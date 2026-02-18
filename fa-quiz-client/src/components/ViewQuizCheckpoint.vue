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
import { computed, onMounted, ref } from "vue";
import { Chart as ChartJS, ArcElement, Tooltip } from "chart.js";
import { Pie } from "vue-chartjs";
import { QuizEntry } from "../quiz";

ChartJS.register(ArcElement, Tooltip);

const props = defineProps<{
    fullData: QuizEntry[];
    chapter: string;
    chapterIndex: number;
    content: string;
    proofAnswer: (selected: number) => void;
    onNext: () => void;
    onBack: () => void;
}>();

// placeholder data
const data = computed(() => {
    const labels = ["Skipped", "Correct", "Incorrect", "Incorrect (even with >2 tries)"];
    const data = [0, 0, 0, 0];
    const backgroundColor = ["#aaa", "#4caf50", "#f44336", "#ff9800", "#9c27b0"];

    console.debug(props.fullData);

    props.fullData.filter(console.log)

    const questions = props.fullData.filter((q) => q.item == "question" && q.chapterIndex == props.chapterIndex) as (QuizEntry & { item: "question" })[];

    console.warn(questions.map(q => ({
        stats: q.stats,
        quest: q.question,
    })))

    for (const q of questions) {
        console.log(q.stats);

        if (q.selectedAnswer === undefined) {
            data[0] += 1;
        } else if (q.stats.triesWrong > 1) {
            data[3] += 1;
        } else if (q.stats.triesWrong === 1) {
            data[2] += 1;
        } else {
            data[1] += 1;
        }
    }

    console.debug(data);

    return { labels, datasets: [{ backgroundColor, data }] };
})
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
