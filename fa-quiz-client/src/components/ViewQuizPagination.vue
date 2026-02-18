<template>
    <div class="view-quiz-pagination-container">
        <div class="view-quiz-pagination">
            <span class="width-filler"></span>
            <span class="question" v-for="question in props.modelValue" :key="question.id">
                <font-awesome-icon v-if="!question.showQuizAnswer" icon="fa-regular fa-circle" class="selected-question" />
                <font-awesome-icon v-else-if="question.selectedAnswer === question.answer" icon="fa-solid fa-circle-check" />
                <font-awesome-icon v-else icon="fa-solid fa-circle-xmark" />
            </span>
            <font-awesome-icon v-for="i in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21]" :key="i" icon="fa-regular fa-circle" />
            <span class="width-filler-sm"></span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { watch } from "vue";

const props = defineProps<{
    modelValue: any[];
}>();

function scrollToActive() {
    // scroll to center the last (active) question
    const container = document.querySelector(".view-quiz-pagination") as HTMLElement;
    if (container && props.modelValue.length > 0) {
        setTimeout(() => {
            const questions = container.querySelectorAll(".question");
            const lastQuestion = questions[questions.length - 1] as HTMLElement;
            
            if (lastQuestion) {
                const questionLeft = lastQuestion.offsetLeft;
                const questionWidth = lastQuestion.offsetWidth;
                const containerWidth = container.offsetWidth;
                const padding = 16; // Match the padding from CSS
                
                // Calculate position to center the question, accounting for padding
                const scrollPosition = questionLeft - (containerWidth / 2) + (questionWidth / 2) - padding;
                
                container.scrollLeft = scrollPosition;
            }
        }, 0);
    }
}

watch(() => props.modelValue, scrollToActive, { deep: true });
</script>

<style lang="scss" scoped>
.view-quiz-pagination-container {
    position: relative;
    margin: auto;
    max-width: 600px;

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
}

.view-quiz-pagination {
    display: flex;
    flex-direction: row;
    gap: 8px;
    padding: 16px;
    overflow-x: auto;
    overflow-y: hidden;
    scroll-behavior: smooth;

    .width-filler {
        flex: 0 0 280px;
    }

    .width-filler-sm {
        flex: 0 0 250px;
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
        width: 24px;
        height: 24px;
        display: inline-block;
        color: #ccc;
    }
}
</style>
