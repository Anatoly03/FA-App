<template>
    <div class="view-quiz-checkpoint">
        <span class="subtitle"> Checkpoint </span>
        <div class="content">
            <div class="badge">
                <span class="data-entry">
                    <span
                        :class="{
                            'green': data.correct > data.total *.6,
                            'yellow': data.correct > data.total *.5 && data.correct <= data.total *.6,
                            'red': data.correct <= data.total *.5
                        }"
                    >{{ Math.round((data.correct / data.total) * 100) }}%</span>
                    <span class="small">Score</span>
                </span>
                <span class="data-entry">
                    <span class="green">{{ data.correct }} / {{ data.total }}</span>
                    <span class="small">Questions</span>
                </span>
            </div>
            <div class="badge badge-border">
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
    const labels = ["Skipped", "Correct", "Incorrect", "Incorrect (even with >2 tries)", "Incorrect (unsolved)"];
    const data = [0, 0, 0, 0, 0];
    const backgroundColor = ["#aaa", "#4caf50", "#f44336", "#c41336", "#000"];

    const questions = props.fullData.filter((q) => q.item == "question" && q.chapterIndex == props.chapterIndex) as (QuizEntry & { item: "question" })[];

    for (const q of questions) {
        if (q.selectedAnswers.length == 0) {
            data[0] += 1;
        } else if (!q.stats.solved) {
            data[4] += 1;
        } else if (q.stats.triesWrong > 1) {
            data[3] += 1;
        } else if (q.stats.triesWrong === 1) {
            data[2] += 1;
        } else {
            data[1] += 1;
        }
    }

    return {
        labels,
        datasets: [{ backgroundColor, data }],
        correct: data[1],
        total: questions.length,
    };
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
    justify-content: center;
    width: 100%;
    gap: 16px;

    .subtitle {
        margin-top: -25px;
        padding: 16px;

        border-bottom: 1px solid #ccc;
    }

    .content {
        display: flex;
        flex-direction: row;
        gap: 8px;
        margin: 0 16px;
    }

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}

.badge {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    width: 30%;
    max-width: 600px;
    gap: 16px;
    padding: 16px;

    &.badge-border {
        border: 1px solid #ccc;
        border-radius: 8px;
    }

    .data-entry {
        display: flex;
        flex-direction: column;
        align-items: center;
        font-size: 3em;
        font-weight: 600;

        .small {
            font-size: 0.5em;
            font-weight: normal;
        }

        .green {
            color: #4caf50;
        }

        .yellow {
            color: #d9ce67;
        }

        .red {
            color: #f44336;
        }
    }
}
</style>
