<template>
    <div class="view-quiz">
        <ViewQuizPagination v-model="previousQuestions" />
        <ViewQuizBody @quiz-fetched="previousQuestions.push($event)" />
    </div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import ViewQuizPagination from "./ViewQuizPagination.vue";
import ViewQuizBody from "./ViewQuizBody.vue";
import { ref } from "vue";

/**
 * Already finished quiz answers will be added to pagination.
 */
const previousQuestions = ref<any[]>([]);
</script>

<style lang="scss" scoped>
.view-quiz {
    display: flex;
    flex-direction: column;
    width: 600px;
    margin: auto;

    transition:
        width 800ms ease,
        height 800ms ease,
        margin 800ms ease;
}
</style>
