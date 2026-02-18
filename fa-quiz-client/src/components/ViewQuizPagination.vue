<template>
    <div class="view-quiz-pagination-container">
        <div class="view-quiz-pagination" ref="containerRef">
            <span class="width-filler"></span>
            <span
                class="quiz-paginationentry"
                v-for="question in props.loadedQuestions"
                :key="question.id"
                :data-question-id="question.id"
                @click="() => props.selectActiveQuestion(question.id)"
            >
                <div class="question" v-if="question.item=='question'">
                    <font-awesome-icon v-if="!question.selectedAnswer" icon="fa-regular fa-circle" class="selected-question" />
                    <font-awesome-icon v-else-if="question.selectedAnswer === question.correctAnswer" icon="fa-solid fa-circle-check" />
                    <font-awesome-icon v-else icon="fa-solid fa-circle-xmark" />
                </div>
                <div class="question" v-if="question.item=='definition'">
                    <font-awesome-icon icon="fa-solid fa-book-open" />
                </div>
            </span>
            <font-awesome-icon v-for="i in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21]" :key="i" icon="fa-regular fa-circle" />
            <span class="width-filler-sm"></span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from "vue";
import { QuizEntry } from "../quiz";

const props = defineProps<{
    loadedQuestions: QuizEntry[];
    activeQuestion: string | null;
    selectActiveQuestion: (questionId: string) => void;
}>();

const containerRef = ref<HTMLElement | null>(null);

function scrollToActive() {
    const container = containerRef.value;
    if (!container || props.loadedQuestions.length === 0) return;

    nextTick(() => {
        let target: HTMLElement | null = null;

        if (props.activeQuestion) {
            target = container.querySelector(
                `[data-question-id="${props.activeQuestion}"]`,
            ) as HTMLElement | null;
        }

        if (!target) {
            const entries = container.querySelectorAll(".quiz-paginationentry");
            target = entries[entries.length - 1] as HTMLElement | null;
        }

        if (!target) return;

        const targetLeft = target.offsetLeft;
        const targetWidth = target.offsetWidth;
        const containerWidth = container.clientWidth;
        const scrollPosition = targetLeft - (containerWidth / 2) + (targetWidth / 2);

        container.scrollTo({ left: scrollPosition, behavior: "smooth" });
    });
}

defineExpose({ scrollToActive });

watch(() => props.loadedQuestions, scrollToActive, { deep: true });
watch(() => props.activeQuestion, scrollToActive);
</script>

<style lang="scss" scoped>
.view-quiz-pagination-container {
    position: relative;
    margin: auto;
    max-width: 600px;
    width: 100%;

    // Gradient overlays at edges
    &::before,
    &::after {
        content: "";
        position: absolute;
        top: 0;
        bottom: 0;
        width: 80px;
        pointer-events: none;
        z-index: 1;
    }

    &::before {
        left: 0;
        background: linear-gradient(to right, #eee, transparent);
    }

    &::after {
        right: 0;
        background: linear-gradient(to left, #eee, transparent);
    }

    @media (max-width: 768px) {
        max-width: 100%;
        
        &::before,
        &::after {
            width: 40px;
        }
    }

    // @media (max-width: 400px) {
    //     display: none;
    // }
}

.view-quiz-pagination {
    display: flex;
    flex-direction: row;
    gap: 8px;
    padding: 16px;
    overflow-x: auto;
    overflow-y: hidden;
    scroll-behavior: smooth;

    @media (max-width: 768px) {
        padding: 12px 8px;
        gap: 6px;
    }

    .width-filler {
        flex: 0 0 calc(50% - 20px);
        
        @media (max-width: 768px) {
            flex: 0 0 calc(50vw - 24px);
        }
    }

    .width-filler-sm {
        flex: 0 0 calc(50% - 12px);
        
        @media (max-width: 768px) {
            flex: 0 0 calc(50vw - 20px);
        }
    }

    .fa-circle,
    .fa-circle-check,
    .fa-circle-xmark {
        width: 24px;
        height: 24px;
        display: inline-block;
        color: #ccc;

        &.selected-question {
            color: black;
        }
        
        @media (max-width: 768px) {
            width: 20px;
            height: 20px;
        }
    }

    .fa-circle-check {
        color: #0a0;
    }

    .fa-circle-xmark {
        color: #a00;
    }

    // Hide scrollbar
    &::-webkit-scrollbar {
        display: none;
    }
    -ms-overflow-style: none;
    scrollbar-width: none;

    .question {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 24px;
        height: 24px;
        // display: inline-block;
        color: black;
        
        @media (max-width: 768px) {
            width: 20px;
            height: 20px;
        }
    }

    .question:has(.fa-book-open) {
        background-color: #ccc;
        border-radius: 50%;
    }
}
</style>
