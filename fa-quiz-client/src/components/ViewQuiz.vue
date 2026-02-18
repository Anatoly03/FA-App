<template>
    <div class="view-quiz">
        <ViewQuizPagination ref="paginationRef" :loadedQuestions="loadedQuestions" :activeQuestion="activeQuizEntry?.id ?? null" :selectActiveQuestion="selectActiveQuestion" />
        <ViewQuizBody :onNext="nextQuestionExists && nextQuizPage" :onBack="activeQuizEntryIndex > 0 && previousQuizPage">
            <ViewQuizQuestion
                v-if="activeQuizEntry && activeQuizEntry.item === 'question'"
                :questionModel="activeQuizEntry"
                :question="activeQuizEntry.question"
                :options="activeQuizEntry.options.map((content, index) => ({ q: index, content }))"
                :answer="activeQuizEntry.correctAnswer"
                :footer="activeQuizEntry.footer"
                :selectedAnswers="activeQuizEntry.selectedAnswers"
                :proofAnswer="proofAnswer"
            />
            <ViewQuizLecture v-else-if="activeQuizEntry && activeQuizEntry.item === 'definition'" :subtitle="activeQuizEntry.subtitle" :content="activeQuizEntry.content" />
            <ViewQuizCheckpoint v-else-if="activeQuizEntry && activeQuizEntry.item === 'chapter-finish'" :chapter="activeQuizEntry.title" :chapterIndex="activeQuizEntry.chapterIndex" :fullData="loadedQuestions" />
        </ViewQuizBody>
    </div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import { computed, onMounted, ref } from "vue";
import ViewQuizPagination from "./ViewQuizPagination.vue";
import ViewQuizBody from "./ViewQuizBody.vue";
import { QuizEntry } from "../quiz";
import { RecordModel } from "pocketbase";
import ViewQuizLecture from "./ViewQuizLecture.vue";
import ViewQuizQuestion from "./ViewQuizQuestion.vue";
import ViewQuizCheckpoint from "./ViewQuizCheckpoint.vue";

/**
 * Current chapter quizzes.
 */
const chapter = ref(0);
const chapterModel = ref<RecordModel | null>(null);

/**
 * Already finished quiz answers will be added to pagination.
 */
const loadedQuestions = ref<QuizEntry[]>([]);
const paginationRef = ref<InstanceType<typeof ViewQuizPagination> | null>(null);
const nextQuestionExists = ref(true);

/**
 * Active quiz entry. This is the quiz that is currently shown in the quiz body.
 */
const activeQuizEntry = computed(() => loadedQuestions.value.find((q) => q.isActive));
const activeQuizEntryIndex = computed(() => loadedQuestions.value.findIndex((q) => q.isActive));

/**
 * Convert a PocketBase question record into a QuizEntry for the quiz.
 */
function getQuizEntryFromQuestion(questionModel: RecordModel, chapterIndex: number) {
    const question = questionModel.question as string;
    const rightOption = questionModel.correctAnswer as string;
    const falseOptions = questionModel.otherAnswers as string[];

    const options = [...falseOptions].sort(() => Math.random() - 0.5);
    const correctAnswer = Math.floor(Math.random() * (options.length + 1));
    options.splice(correctAnswer, 0, rightOption);

    return {
        item: "question",
        id: questionModel.id,
        chapterIndex,
        question,
        options: options.map((content, index) => content),
        footer: questionModel.footer,
        correctAnswer,
        showQuizAnswer: false,
        selectedAnswers: [],
        isActive: false,
        stats: {
            solved: false,
            tries: 0,
            triesWrong: 0,
        },
    } as QuizEntry;
}

/**
 * Convert a PocketBase lecture record into a QuizEntry for the quiz.
 */
function getQuizEntryFromLecture(lectureModel: RecordModel) {
    const subtitle = lectureModel.subtitle as string;
    const content = lectureModel.content as string;

    return {
        item: "definition",
        id: lectureModel.id,
        isActive: false,
        subtitle,
        content,
    } as QuizEntry;
}

/**
 * Convert a PocketBase lecture record into a QuizEntry for the quiz.
 */
function getQuizEntryFromChapterFinish(chapterModel: RecordModel, chapterIndex: number) {
    const title = chapterModel.title as string;

    return {
        item: "chapter-finish",
        id: chapterModel.id,
        isActive: false,
        title,
        chapterIndex,
    } as QuizEntry;
}

/**
 * Re-scroll pagination to active quiz entry.
 */
function updatePagination() {
    nextQuestionExists.value = true;
    paginationRef.value?.scrollToActive();
}

/**
 * Advance pagination. Find the current active question and set the next one
 * as active. Additionally control chapter locks.
 */
async function nextQuizPage() {
    if (!nextQuestionExists.value) return;

    const activeIndex = loadedQuestions.value.findIndex((q) => q.isActive);
    if (activeIndex === -1) activeIndex = loadedQuestions.value.length - 1;

    if (activeIndex < 0) throw new Error("unreachable");

    if (activeIndex < loadedQuestions.value.length - 1) {
        // deactivate
        loadedQuestions.value[activeIndex].isActive = false;

        // activate next
        loadedQuestions.value[activeIndex + 1].isActive = true;
        return updatePagination();
    }

    // if no next quiz, load next chapter
    chapter.value += 1;
    await fetchChapter(false);

    nextQuizPage();
}

/**
 * Go back in pagination. Find the current active question and set the previous
 * one as active. Additionally control chapter locks.
 */
async function previousQuizPage() {
    const activeIndex = loadedQuestions.value.findIndex((q) => q.isActive);
    if (activeIndex === -1) throw new Error("unreachable");

    if (activeIndex > 0) {
        // deactivate
        loadedQuestions.value[activeIndex].isActive = false;

        // activate previous
        loadedQuestions.value[activeIndex - 1].isActive = true;
        return updatePagination();
    }

    // if no previous quiz, go to previous chapter
    chapter.value -= 1;
    await fetchChapter(false);

    previousQuizPage();
}

function selectActiveQuestion(id: string) {
    for (const question of loadedQuestions.value) question.isActive = question.id === id;
    return updatePagination();
}

/**
 * Proof the answer selected by the user. Set the selected answer and show the
 * quiz answer.
 */
function proofAnswer(selected: number) {
    if (!activeQuizEntry.value) return;

    console.warn('Selected answer', selected, 'for question', activeQuizEntry.value.question);

    activeQuizEntry.value.selectedAnswers.push(selected);
    activeQuizEntry.value.showQuizAnswer = true;

    // record statistics
    activeQuizEntry.value.stats.tries += 1;
    if (selected === activeQuizEntry.value.correctAnswer) {
        activeQuizEntry.value.stats.solved = true;
    } else {
        activeQuizEntry.value.stats.triesWrong += 1;
    }
}

/**
 * Fetch the next quiz from the API and set it to the `quiz` variable.
 */
async function fetchChapter(resetActive = true) {
    try {
        // fetch all lectures for the current chapter
        const chapterLectures = await pb.collection("chapters").getFirstListItem(`index = ${chapter.value}`, {
            requestKey: "chapter-" + chapter.value,
            expand: "lessons",
        });
        chapterModel.value = chapterLectures;

        // fetch all quizzes for the current chapter
        const chapterQuestions = await pb.collection("mc_questions").getFullList({
            requestKey: "questions-chapter-" + chapter.value,
            expand: "chapter",
            filter: `chapter.index = ${chapter.value}`,
        });

        // load lectures
        if (chapterLectures.expand.lessons && chapterLectures.expand.lessons.length > 0) loadedQuestions.value.push(...chapterLectures.expand.lessons.map(getQuizEntryFromLecture));

        // load questions
        const shuffled = chapterQuestions.map(q => getQuizEntryFromQuestion(q, chapter.value)).sort(() => Math.random() - 0.5);
        loadedQuestions.value.push(...shuffled);

        // load checkpoint (new chapter announce)
        loadedQuestions.value.push(getQuizEntryFromChapterFinish(chapterLectures, chapter.value));

        // select first
        if (resetActive) {
            loadedQuestions.value[0].isActive = true;
            updatePagination();
        }
    } catch (e) {
        console.error("Error fetching chapter", e);
        nextQuestionExists.value = false;
        return;
    }
}

onMounted(() => {
    fetchChapter(true);
});
</script>

<style lang="scss" scoped>
.view-quiz {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: auto;

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}
</style>
