<template>
    <div class="view-quiz-body">
        <span class="triangle-pointer"></span>
        <slot></slot>
        <div class="pagination">
            <button @click="props.onBack()" :disabled="!props.onBack">
                <font-awesome-icon icon="fa-regular fa-circle-left" />
            </button>
            <button @click="props.onNext()" :disabled="!props.onNext">
                <font-awesome-icon icon="fa-regular fa-circle-right" />
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
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
    top: -12px;

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
        padding: 16px;

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
        width 1800ms ease,
        height 1800ms ease,
        margin 1800ms ease;
}
</style>
