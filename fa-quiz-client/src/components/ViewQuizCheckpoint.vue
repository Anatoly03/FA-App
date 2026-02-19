<template>
    <div class="view-quiz-checkpoint">
        <span class="subtitle"> Checkpoint </span>
        <div class="content">
            <div class="badge badge-left">
                <div class="badge-stats">
                    <span class="data-entry">
                        <span
                            :class="{
                                green: data.correct > data.total * 0.6,
                                yellow: data.correct > data.total * 0.5 && data.correct <= data.total * 0.6,
                                red: data.correct <= data.total * 0.5,
                            }"
                            >{{ Math.round((data.correct / data.total) * 100) }}%</span
                        >
                        <span class="small">Score</span>
                    </span>
                    <span class="data-entry">
                        <span class="green">{{ data.correct }} / {{ data.total }}</span>
                        <span class="small">Questions</span>
                    </span>
                </div>
                <div class="horizontal">
                    <a class="btn" @click="resetLecture('all')">Reset Lecture</a>
                    <a class="btn" @click="resetLecture('wrong')">Reset Wrong</a>
                </div>
            </div>
            <div class="badge badge-border">
                <div class="chart-wrap">
                    <Pie :data="data" :options="{ onClick }" />
                </div>
                <span class="chapter">{{ props.chapter }}</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { Chart as ChartJS, ArcElement, Tooltip, type ActiveElement, type ChartEvent } from "chart.js";
import { Pie } from "vue-chartjs";
import { QuizEntry } from "../quiz";

ChartJS.register(ArcElement, Tooltip);

const props = defineProps<{
    fullData: QuizEntry[];
    chapter: string;
    chapterIndex: number;
    scrollBackTo: (lambda: (entry: QuizEntry) => boolean) => void;
    scrollToPreviousCheckpoint: (kind: "all" | "wrong", chapterIndex: number) => void;
}>();

// placeholder data
const data = computed(() => {
    const labels = ["Skipped", "Correct", "Incorrect", "Incorrect (even with >2 tries)", "Incorrect (unsolved)"];
    // const firstQuestions: (string | null)[] = labels.map(() => null)
    const data = [0, 0, 0, 0, 0];
    const backgroundColor = ["#aaa", "#4caf50", "#f44336", "#c41336", "#000"];

    const questions = props.fullData.filter((q) => q.item == "question" && q.chapterIndex == props.chapterIndex) as (QuizEntry & { item: "question" })[];

    for (const q of questions) {
        if (q.selectedAnswers.length == 0) {
            data[0] += 1;
            // firstQuestions[0] ??= q.question;
        } else if (!q.stats.solved) {
            data[4] += 1;
            // firstQuestions[4] ??= q.question;
        } else if (q.stats.triesWrong > 1) {
            data[3] += 1;
            // firstQuestions[3] ??= q.question;
        } else if (q.stats.triesWrong === 1) {
            data[2] += 1;
            // firstQuestions[2] ??= q.question;
        } else {
            data[1] += 1;
            // firstQuestions[1] ??= q.question;
        }
    }

    // for (let i = 0; i < data.length; i++) {
    //     if (data[i] == 1) {
    //         labels[i] += `: ${firstQuestions[i]}`;
    //     }
    // }

    return {
        labels,
        datasets: [{ backgroundColor, data }],
        correct: data[1],
        total: questions.length,
    };
});

function resetLecture(kind: "all" | "wrong") {
    props.scrollToPreviousCheckpoint(kind, props.chapterIndex);
}

function onClick(_event: ChartEvent, elements: ActiveElement[]) {
    if (elements.length === 0) return;

    const _type = elements[0].datasetIndex;
    const lambda = (entry: QuizEntry): boolean => {
        if (entry.item !== "question") return false;
        switch (_type) {
            case 0: return entry.selectedAnswers.length == 0;
            case 1: return entry.stats.solved && entry.stats.triesWrong === 0;
            case 2: return !entry.stats.solved && entry.stats.triesWrong === 1;
            case 3: return !entry.stats.solved && entry.stats.triesWrong > 1;
            case 4: return !entry.stats.solved && entry.stats.triesWrong === 0;
            default: return false;
        }
    };

    props.scrollBackTo(lambda);
}
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
    padding: 8px 8px 0;
    box-sizing: border-box;
    margin-bottom: -8px;

    .subtitle {
        margin-top: -25px;
        padding: 16px;

        border-bottom: 1px solid #ccc;
    }

    .content {
        display: grid;
        grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
        align-items: stretch;
        gap: 8px;
        margin: 0 8px;

        @media (max-width: 768px) {
            grid-template-columns: 1fr;
            margin: 0;
        }
    }

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}

.badge {
    flex: 1;
    min-width: 0;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    width: auto;
    max-width: none;
    gap: 16px;

    @media (max-width: 768px) {
        max-width: none;
    }

    &:first-child {
        flex: 1 1 0;
    }

    .badge-stats {
        padding: 16px;
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
    }

    .horizontal {
        width: 100%;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        gap: 4px;

        .btn {
            align-items: center;
            text-align: center;
            white-space: nowrap;
            line-height: 1.1;
            font-size: 1.05em;

            @media (max-width: 768px) {
                font-size: 1.1em;
            }
        }
    }

    .badge-left {
        flex: 1;
    }

    .btn {
        display: flex;
        flex: 1;
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

        display: flex;
        padding: 8px;
        font-size: 1.2em;
        justify-content: center;

        &.long-btn {
            flex: 1;
        }
        
        @media (max-width: 768px) {
            padding: 12px;
            font-size: 1.5em;
            min-height: 44px;
        }
    }

    &.badge-border {
        flex: 1 1 0;
        padding: 16px;
        border: 1px solid #ccc;
        border-radius: 8px;
        justify-content: flex-start;
        align-items: center;

        .chart-wrap {
            width: min(100%, 240px);
            margin: 0 auto;
        }

        :deep(canvas) {
            width: 100% !important;
            height: auto !important;
            display: block;
        }

        .chapter {
            width: 100%;
            text-align: center;
            font-size: 1.05em;
            line-height: 1.25;
            overflow-wrap: anywhere;
        }
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
