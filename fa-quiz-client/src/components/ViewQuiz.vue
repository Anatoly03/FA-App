<template>
    <div class="view-quiz">
        <ViewQuizPagination ref="paginationRef" :loadedQuestions="loadedQuestions" :activeQuestion="activeQuizEntry?.id ?? null" :selectActiveQuestion="selectActiveQuestion" />
        <ViewQuizBody
            v-if="activeQuizEntry && activeQuizEntry.item === 'question'"
            :question="activeQuizEntry.question"
            :options="activeQuizEntry.options.map((content, index) => ({ q: index, content }))"
            :answer="activeQuizEntry.correctAnswer"
            :footer="activeQuizEntry.footer"
            :selectedAnswer="activeQuizEntry.selectedAnswer"
            :proofAnswer="proofAnswer"
            :onNext="nextQuizPage"
        />
        <ViewQuizLecture v-else-if="activeQuizEntry && activeQuizEntry.item === 'definition'" :subtitle="activeQuizEntry.subtitle" :content="activeQuizEntry.content" :onNext="nextQuizPage" />
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

/**
 * Current chapter quizzes.
 */
const chapter = ref(0);

/**
 * Already finished quiz answers will be added to pagination.
 */
const loadedQuestions = ref<QuizEntry[]>([]);
const paginationRef = ref<InstanceType<typeof ViewQuizPagination> | null>(null);

/**
 * Active quiz entry. This is the quiz that is currently shown in the quiz body.
 */
const activeQuizEntry = computed(() => loadedQuestions.value.find((q) => q.isActive));

/**
 * Convert a PocketBase question record into a QuizEntry for the quiz.
 */
function getQuizEntryFromQuestion(questionModel: RecordModel) {
    const question = questionModel.question as string;
    const rightOption = questionModel.correctAnswer as string;
    const falseOptions = questionModel.otherAnswers as string[];

    const options = [...falseOptions].sort(() => Math.random() - 0.5);
    const correctAnswer = Math.floor(Math.random() * (options.length + 1));
    options.splice(correctAnswer, 0, rightOption);

    return {
        item: "question",
        id: questionModel.id,
        question,
        options: options.map((content, index) => content),
        footer: questionModel.footer,
        correctAnswer,
        showQuizAnswer: false,
        selectedAnswer: undefined,
        isActive: false,
    } as QuizEntry;
}

/**
 * Convert a PocketBase lecture record into a QuizEntry for the quiz.
 */
function getQuizEntryFromLecture(lectureModel: RecordModel) {
    const subtitle = lectureModel.subtitle as string;
    const content = lectureModel.content as string;

    console.debug(lectureModel);

    return {
        item: "definition",
        id: lectureModel.id,
        isActive: false,
        subtitle,
        content,
    } as QuizEntry;
}

/**
 * Re-scroll pagination to active quiz entry.
 */
function updatePagination() {
    paginationRef.value?.scrollToActive();
}

/**
 * Advance pagination. Find the current active question and set the next one
 * as active. Additionally control chapter locks.
 */
async function nextQuizPage() {
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

    activeQuizEntry.value.selectedAnswer = selected;
    activeQuizEntry.value.showQuizAnswer = true;
}

/**
 * Fetch the next quiz from the API and set it to the `quiz` variable.
 */
async function fetchChapter(resetActive = true) {
    // fetch all lectures for the current chapter
    const chapterLectures = await pb.collection("chapters").getFirstListItem(`index = ${chapter.value}`, {
        requestKey: "chapter-" + chapter.value,
        expand: "lessons",
    });

    // fetch all quizzes for the current chapter
    const chapterQuestions = await pb.collection("mc_questions").getFullList({
        requestKey: "questions-chapter-" + chapter.value,
        expand: "chapter",
        filter: `chapter.index = ${chapter.value}`,
    });

    // load lectures
    if (chapterLectures.expand.lessons && chapterLectures.expand.lessons.length > 0)
        loadedQuestions.value.push(...chapterLectures.expand.lessons.map(getQuizEntryFromLecture));

    // load questions
    const shuffled = chapterQuestions.map(getQuizEntryFromQuestion).sort(() => Math.random() - 0.5);
    loadedQuestions.value.push(...shuffled);

    // select first
    if (resetActive) {
        loadedQuestions.value[0].isActive = true;
        updatePagination();
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
