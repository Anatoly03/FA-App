<template>
    <div class="view-quiz-lecture">
        <span class="subtitle" v-html="props.subtitle"></span>
        <div class="content" v-html="props.content"></div>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    subtitle: string;
    content: string;
    proofAnswer: (selected: number) => void;
    onNext: () => void;
    onBack: () => void;
}>();

const showQuizAnswer = computed(() => props.selectedAnswers.length > 0);
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

.github-link {
    border-radius: 6px;
    display: flex;
    background-color: #469b46;
    padding: 10px;
    text-decoration: none;
    color: white;

    &:hover {
        background-color: #70bb70;
    }
}

p:has(.gitub-link) {
    display: flex;
    width: 100%;
    flex-direction: column;
}
</style>

<style lang="scss" scoped>
.view-quiz-lecture {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 16px;

    .subtitle {
        margin-top: -25px;
        padding: 16px;

        border-bottom: 1px solid #ccc;
    }

    .content {
        padding: 16px;
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
        padding: 16px;
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
        padding: 16px;
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
