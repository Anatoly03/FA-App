<template>
    <div class="view-quiz-pagination-container">
        <div class="view-quiz-pagination" ref="containerRef">
            <span class="width-filler"></span>
            <span
                class="quiz-pagination-entry"
                v-for="question in props.loadedQuestions"
                :key="question.id"
                :data-question-id="question.id"
                role="button"
                @click.stop="() => props.selectActiveQuestion(question.id)"
                @keydown.enter.stop="() => props.selectActiveQuestion(question.id)"
                @keydown.space.stop="() => props.selectActiveQuestion(question.id)"
            >
                <div class="question" v-if="question.item=='question'">
                    <font-awesome-icon icon="fa-solid fa-face fa-face-laugh-beam" v-if="question.isActive && question.stats.solved && question.stats.triesWrong === 0" />
                    <font-awesome-icon icon="fa-solid fa-face fa-face-meh" v-else-if="question.isActive && question.stats.triesWrong == 1" />
                    <font-awesome-icon icon="fa-solid fa-face fa-face-frown" v-else-if="question.isActive && question.stats.triesWrong > 1" />
                    <font-awesome-icon icon="fa-regular fa-circle-dot" v-else-if="question.isActive" />
                    <font-awesome-icon icon="fa-regular fa-circle" class="question" v-else-if="!question.stats.solved && question.stats.triesWrong === 0" />
                    <font-awesome-icon icon="fa-solid fa-circle-check" v-else-if="question.stats.triesWrong === 0" />
                    <font-awesome-icon icon="fa-solid fa-circle-xmark" v-else />
                </div>
                <div class="question" v-if="question.item=='definition'">
                    <font-awesome-icon icon="fa-regular fa-compass" />
                </div>
                <div class="question" v-if="question.item=='chapter-finish'">
                    <font-awesome-icon icon="fa-solid fa-lock" />
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
        width: 240px;
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
        pointer-events: none;

        &.selected-question {
            color: black;
        }

        @media (max-width: 768px) {
            width: 20px;
            height: 20px;
        }
    }

    .fa-circle-check,
    .fa-face-laugh-beam {
        color: #0a0;
    }

    .fa-face-meh,
    .fa-face-frown,
    .fa-circle-xmark {
        color: #ff4000;
    }

    .fa-cube {
        font-size: 1.6em;
        pointer-events: none;
    }

    .fa-compass {
        font-size: 1.5em;
        pointer-events: none;
    }

    .fa-face-laugh-beam,
    .fa-face-meh,
    .fa-face-frown,
    .fa-circle-dot {
        font-size: 1.5em;
    }

    .quiz-paginationentry {
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        flex-shrink: 0;
        padding: 4px;
        pointer-events: auto;
        user-select: none;
        transition: transform 150ms ease, opacity 150ms ease;
        
        &:hover {
            transform: scale(1.15);
            opacity: 0.8;
        }
        
        &:active {
            transform: scale(0.95);
        }
        
        &:focus-visible {
            outline: 2px solid #333;
            outline-offset: 2px;
            border-radius: 50%;
        }
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
        color: black;
        pointer-events: none;
        
        @media (max-width: 768px) {
            width: 20px;
            height: 20px;
        }
    }

    // .question:has(.fa-book-open) {
    //     background-color: #ccc;
    //     border-radius: 50%;
    // }
}
</style>
