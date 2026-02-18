<template>
    <div class="view-quiz-body">
        <span class="triangle-pointer"></span>
        <span v-html="props.question"></span>
        <div class="options">
            <button
                v-for="question in props.options"
                :key="question.q + question.content"
                @click="props.proofAnswer(question.q)"
                :class="{
                    correct: showQuizAnswer && question.q === props.answer,
                    incorrect: showQuizAnswer && question.q !== props.answer && props.selectedAnswer === question.q,
                }"
                :disabled="showQuizAnswer"
            >
                <span class="option-selection-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-dot" v-if="showQuizAnswer && question.q == props.selectedAnswer" />
                    <font-awesome-icon icon="fa-regular fa-circle" v-if="!showQuizAnswer || question.q != props.selectedAnswer" />
                </span>
                <span class="option-item">
                    {{ question.content }}
                </span>
                <span class="option-correction-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-check" v-if="showQuizAnswer && question.q === props.answer" />
                    <font-awesome-icon icon="fa-regular fa-circle-xmark" v-else-if="showQuizAnswer && question.q !== props.answer" />
                </span> 
            </button>
        </div>
        <span class="footer" v-if="props.footer">{{ props.footer }}</span>
        <div class="pagination">
            <button @click="props.onBack()" :disabled="!props.onBack">
                <font-awesome-icon icon="fa-regular fa-circle-left" />
            </button>
            <button @click="props.onNext()" :disabled="!props.onNext || !showQuizAnswer">
                <font-awesome-icon icon="fa-regular fa-circle-right" />
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    question: string;
    options: { q: number; content: string }[];
    answer: number;
    footer: string | null;
    selectedAnswer: number | undefined;

    proofAnswer: (selected: number) => void;
    onNext: () => void;
    onBack: () => void;
}>();

const showQuizAnswer = computed(() => props.selectedAnswer !== undefined);

// const allQuizzes = ref<{}[]>([]);

// const emit = defineEmits<{
//     (e: "quiz-next", quiz: any): void;
//     (e: "quiz-fetched", quiz: any): void;
// }>();

// const quiz = ref({
//     id: 'unknown',
//     question: "Q?",
//     options: [
//         { q: 1, content: "Option A" },
//         { q: 2, content: "Option B" },
//         { q: 3, content: "Option C" },
//         { q: 4, content: "Option D" },
//     ],
//     footer: null,
//     answer: 2,
//     showQuizAnswer: false,
//     selectedAnswer: undefined as number | undefined,
// });

// /**
//  * Fetch a quiz from the API and set it to the `quiz` variable.
//  */
// async function fetchQuiz() {
//     const response = await pb.collection("mc_questions").getFullList();
//     allQuizzes.value = response;

//     // const singleResponse = await pb.collection("mc_questions").getOne('okdaal45c1rpzcb');
//     // allQuizzes.value = [singleResponse];

//     showQuiz();
// }

// /**
//  * Show random quiz from the list of quizzes fetched from the API.
//  */
// function showQuiz() {
//     quiz.value.showQuizAnswer = false;

//     const randomIndex = Math.floor(Math.random() * allQuizzes.value.length);
//     const randomQuiz: any = allQuizzes.value[randomIndex];

//     // deep clone of other answers so splice() does not mutate database
//     const otherAnswers = [...randomQuiz.otherAnswers]
//         .sort(() => Math.random() - 0.5);
//     const correctAnswer = randomQuiz.correctAnswer;

//     // add correct answer to random position in other answers
//     // and track position
//     const randomInsert = Math.floor(Math.random() * (otherAnswers.length + 1));
//     otherAnswers.splice(randomInsert, 0, correctAnswer);

//     quiz.value.id = randomQuiz.id;
//     quiz.value.question = formatQuestion(randomQuiz.question);
//     quiz.value.options = otherAnswers.map((content: string, q: number) => ({ q, content }));
//     quiz.value.answer = randomInsert;
//     quiz.value.footer = randomQuiz.footer;
    
//     // Emit quiz-next event when moving to next question
//     emit("quiz-next", JSON.parse(JSON.stringify(quiz.value)));
// }

// /**
//  * Transform the question highlighting keywords (e.g. 'not') into HTML with <b> tags.
//  */
// function formatQuestion(question: string) {
//     const keywords = ["not", "cannot", "correct", "primary"];
//     let formatted = question;
//     for (const keyword of keywords) {
//         const regex = new RegExp(`\\b${keyword}\\b`, "gi");
//         formatted = formatted.replace(regex, `<mark>${keyword}</mark>`);
//     }
//     return formatted;
// }

// /**
//  * Proof answer
//  */
// function proofAnswer(selected: number) {
//     quiz.value.showQuizAnswer = true;
//     quiz.value.selectedAnswer = selected;

//     // emit quiz for pagination
//     emit("quiz-fetched", JSON.parse(JSON.stringify(quiz.value)));
// }

// // fetch a quiz on component mount
// onBeforeMount(async () => await fetchQuiz());
</script>

<style lang="scss">
mark {
    // padding: 2px;
    background-color: inherit;
    // background-color: #ffeb37;
    // color: #000;
    font-weight: 600;
    text-decoration: underline;
}
</style>

<style lang="scss" scoped>
.triangle-pointer {
    position: relative;
    top: -28px;

    width: 0;
    height: 0;
    border-left: 12px solid transparent;
    border-right: 12px solid transparent;
    border-bottom: 12px solid #bbb;
    margin: auto;

    @media (max-width: 768px) {
        top: -24px;
    }
}

.view-quiz-body {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 600px;
    margin: auto;
    gap: 16px;

    border: 1px solid #ccc;
    border-radius: 8px;
    padding: 16px;

    @media (max-width: 768px) {
        width: 80%;
        margin: auto;
        padding: 12px;
        gap: 12px;
        border-radius: 6px;
        font-size: 16px; // Prevent iOS text zoom
    }

    .options {
        display: flex;
        flex-direction: column;
        width: 100%;
        margin: auto;
        gap: 4px;

        @media (max-width: 768px) {
            gap: 8px;
        }

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

        @media (max-width: 768px) {
            padding: 8px;
            min-height: 44px; // Minimum touch target size
            align-items: center;
        }

        &:hover:not([disabled]) {
            background-color: #ddd;
        }

        &.correct {
            background-color: #c8e6c9;

            &:hover:not([disabled]) {
                background-color: #b7d5b8;
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

    .footer {
        color: #888;
    }

    .pagination {
        display: flex;
        flex-direction: row;
        gap: 8px;
        justify-content: space-between;

        button {
            display: flex;
            padding: 8px;
            font-size: 1.2em;
            flex: 0.4;
            justify-content: center;
            
            @media (max-width: 768px) {
                padding: 12px;
                font-size: 1.5em;
                min-height: 44px;
            }
        }
    }

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}
</style>
