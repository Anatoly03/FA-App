<template>
    <div class="view-quiz-checkpoint">
        <span class="subtitle"> Checkpoint </span>
        <div class="content">
            <div class="checkpoint-badge">
                <Pie :data="data" />
                {{ props.chapter }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { Chart as ChartJS, ArcElement, Tooltip } from "chart.js";
import { Pie } from "vue-chartjs";

ChartJS.register(ArcElement, Tooltip);

const data = {
    labels: ["VueJs", "EmberJs", "ReactJs", "AngularJs"],
    datasets: [
        {
            backgroundColor: ["#41B883", "#E46651", "#00D8FF", "#DD1B16"],
            data: [40, 20, 80, 10],
        },
    ],
};

const props = defineProps<{
    chapter: string;
    content: string;
    proofAnswer: (selected: number) => void;
    onNext: () => void;
    onBack: () => void;
}>();

const showQuizAnswer = computed(() => props.selectedAnswer !== undefined);
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
.view-quiz-checkpoint {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 16px;

    .subtitle {
        margin-top: -25px;
        padding: 16px;

        border-bottom: 1px solid #ccc;
    }

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}

.checkpoint-badge {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 30%;
    max-width: 600px;
    margin: auto;
    gap: 16px;

    border: 1px solid #ccc;
    border-radius: 8px;
    padding: 16px;
}
</style>
