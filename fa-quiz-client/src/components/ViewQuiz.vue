<template>
    <div class="view-quiz">
        <span v-html="quiz.question"></span>
        <div class="options">
            <a
                v-for="question in quiz.options"
                :key="question.q + question.content"
                @click="proofAnswer(question.q)"
                :class="{
                    correct: showQuizAnswer && question.q === quiz.answer,
                    incorrect: showQuizAnswer && question.q !== quiz.answer && selectedAnswer === question.q,
                }"
            >
                <span class="option-selection-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-dot" v-if="showQuizAnswer && question.q == selectedAnswer" />
                    <font-awesome-icon icon="fa-regular fa-circle" v-if="!showQuizAnswer || question.q != selectedAnswer" />
                </span>
                <span class="option-item">
                    {{ question.content }}
                </span>
                <span class="option-correction-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-check" v-if="showQuizAnswer && question.q === quiz.answer" />
                    <font-awesome-icon icon="fa-regular fa-circle-xmark" v-else-if="showQuizAnswer && question.q !== quiz.answer" />
                </span> 
            </a>
        </div>
        <div class="pagination">
            <button @click="showQuiz()" disabled>Back</button>
            <button @click="showQuiz()" :disabled="!showQuizAnswer">Next</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import { onBeforeMount, onMounted, ref } from "vue";

const showQuizAnswer = ref(false);
const selectedAnswer = ref<number | null>(null);

const allQuizzes = ref<{}[]>([]);

const quiz = ref({
    question: "Q?",
    options: [
        { q: 1, content: "Option A" },
        { q: 2, content: "Option B" },
        { q: 3, content: "Option C" },
        { q: 4, content: "Option D" },
    ],
    answer: 2,
});

/**
 * Fetch a quiz from the API and set it to the `quiz` variable.
 */
async function fetchQuiz() {
    const response = await pb.collection("mc_questions").getFullList();
    allQuizzes.value = response;

    showQuiz();
}

/**
 * Show random quiz from the list of quizzes fetched from the API.
 */
function showQuiz() {
    showQuizAnswer.value = false;

    const randomIndex = Math.floor(Math.random() * allQuizzes.value.length);
    const randomQuiz: any = allQuizzes.value[randomIndex];

    // deep clone of other answers so splice() does not mutate database
    const otherAnswers = [...randomQuiz.otherAnswers]
        .sort(() => Math.random() - 0.5);
    const correctAnswer = randomQuiz.correctAnswer;

    // add correct answer to random position in other answers
    // and track position
    const randomInsert = Math.floor(Math.random() * (otherAnswers.length + 1));
    otherAnswers.splice(randomInsert, 0, correctAnswer);

    quiz.value.question = randomQuiz.question;
    quiz.value.options = otherAnswers.map((content: string, q: number) => ({ q, content }));
    quiz.value.answer = randomInsert;
}

/**
 * Proof answer
 */
function proofAnswer(selected: number) {
    showQuizAnswer.value = true;
    selectedAnswer.value = selected;
}

// fetch a quiz on component mount
onBeforeMount(async () => await fetchQuiz());
</script>

<style lang="scss" scoped>
.view-quiz {
    display: flex;
    flex-direction: column;
    width: 600px;
    margin: auto;
    gap: 16px;

    border: 1px solid #ccc;
    border-radius: 8px;
    padding: 16px;

    .options {
        display: flex;
        flex-direction: column;
        width: 100%;
        margin: auto;
        gap: 4px;

        .option-item {
            display: flex;
            flex: 1;
            padding: 4px;
            border-radius: 4px;
        }

        .option-correction-marker, .option-selection-marker {
            display: flex;
            width: 24px;
            justify-content: center;
            align-items: center;

            .fa-circle-check {
                color: green;
            }

            .fa-circle-xmark {
                color: darkred;
            }
        }
    }

    a,
    button {
        display: flex;
        border: 1px solid #ccc;
        border-radius: 4px;
        padding: 4px;
        background-color: #eee;

        &:hover:not([disabled]) {
            background-color: #ddd;
        }

        &.correct {
            background-color: #c8e6c9;

            &:hover:not([disabled]) {
                background-color: #f6f9f6;
            }
        }

        &.incorrect {
            background-color: #ffcdd2;

            &:hover:not([disabled]) {
                background-color: #eebcc1;
            }
        }

        transition: background-color 160ms ease;
    }

    .pagination {
        display: flex;
        flex-direction: row;
        gap: 8px;
        justify-content: space-between;

        button {
            display: flex;
            flex: 0.4;
            justify-content: center;
        }
    }

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}
</style>
