<template>
    <div class="view-quiz">
        <ViewQuizPagination
            ref="paginationRef"
            :loadedQuestions="loadedQuestions"
            :activeQuestion="activeQuizEntry?.id ?? null"
            :selectActiveQuestion="selectActiveQuestion"
        />
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
    </div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import { computed, onMounted, ref } from "vue";
import ViewQuizPagination from "./ViewQuizPagination.vue";
import ViewQuizBody from "./ViewQuizBody.vue";
import { QuizEntry } from "../quiz";
import { RecordModel } from "pocketbase";

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
const activeQuizEntry = computed(() => loadedQuestions.value.find(q => q.isActive));

/**
 * Convert a PocketBase question record into a QuizEntry for the quiz.
 */
function getQuizEntryFromQuestion(questionModel: RecordModel) {
    const question = questionModel.question as string;
    const rightOption = questionModel.rightAnswer as string;
    const falseOptions = questionModel.otherAnswers as string[];

    const options = [...falseOptions]
            .sort(() => Math.random() - 0.5);
    const correctAnswer = Math.floor(Math.random() * (options.length + 1));
    options.splice(correctAnswer, 0, rightOption);
            
    return {
        item: 'question',
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

function updatePagination() {
    paginationRef.value?.scrollToActive();
}

/**
 * Advance pagination. Find the current active question and set the next one
 * as active. Additionally control chapter locks.
 */
function nextQuizPage() {
    const activeIndex = loadedQuestions.value.findIndex(q => q.isActive);
    if (activeIndex === -1) activeIndex = loadedQuestions.value.length - 1;

    if (activeIndex < 0) throw new Error('unreachable')

    // deactivate
    loadedQuestions.value[activeIndex].isActive = false;

    // activate next
    if (activeIndex < loadedQuestions.value.length - 1) {
        loadedQuestions.value[activeIndex + 1].isActive = true;
        return updatePagination();
    }
}

function selectActiveQuestion(id: string) {
    for (const question of loadedQuestions.value)
        question.isActive = question.id === id;
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
async function fetchChapter() {
    // fetch all quizzes for the current chapter
    const chapterQuestions = await pb.collection("mc_questions").getFullList({
        requestKey: 'chapter-' + chapter.value,
        expand: "chapter",
        filter: `chapter.index = ${chapter.value}`,
    });

    // load questions
    for (const question of chapterQuestions) {
        loadedQuestions.value.push(getQuizEntryFromQuestion(question));
    }

    // select first
    loadedQuestions.value[0].isActive = true;
    updatePagination();
}

onMounted(() => {
    fetchChapter();
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
