<template>
    <div class="view-quiz-question">
        <span class="question" v-html="highlightWords(props.question)"></span>
        <div class="options">
            <button
                v-for="question in props.options"
                :key="question.q + question.content"
                @click="props.proofAnswer(question.q)"
                :class="{
                    correct: props.questionModel.selectedAnswers.includes(question.q) && question.q === props.answer,
                    incorrect: props.questionModel.selectedAnswers.includes(question.q) && question.q !== props.answer,
                }"
                :disabled="props.questionModel.stats.solved"
            >
                <span class="option-selection-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-dot" v-if="props.questionModel.selectedAnswers.includes(question.q)" />
                    <font-awesome-icon icon="fa-regular fa-circle" v-else />
                </span>
                <span class="option-item">
                    {{ question.content }}
                </span>
                <span class="option-correction-marker">
                    <font-awesome-icon icon="fa-regular fa-circle-check" v-if="props.questionModel.selectedAnswers.includes(question.q) && question.q === props.answer" />
                    <font-awesome-icon icon="fa-regular fa-circle-xmark" v-else-if="props.questionModel.selectedAnswers.includes(question.q) && question.q !== props.answer" />
                </span> 
            </button>
        </div>
        <span class="footer" v-if="props.footer">{{ props.footer }}</span>
    </div>
</template>

<script setup lang="ts">
import { QuizEntry } from "../quiz";

function highlightWords(text: string) {
    const wordsToHighlight = ["not", "cannot", "primary", "correct", "incorrect"];
    let highlightedText = text;

    for (const word of wordsToHighlight) {
        const regex = new RegExp(`(${word})`, "gi");
        highlightedText = highlightedText.replace(regex, `<mark>$1</mark>`);
    }

    return highlightedText;
}

const props = defineProps<{
    questionModel: QuizEntry & { item: "question" };
    question: string;
    options: { q: number; content: string }[];
    answer: number;
    footer: string | null;

    proofAnswer: (selected: number) => void;
}>();
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
.view-quiz-question {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin: auto;
    gap: 16px;

    .question {
        margin-top: -25px;
        padding: 16px;

        border-bottom: 1px solid #ccc;
    }

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
        gap: 4px;
        margin: 0;
        padding: 16px;
        box-sizing: border-box;

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

        button{
            display: flex;
            width: 100%;
            box-sizing: border-box;
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
