<template>
    <div class='app-body'>
        <router-view></router-view>
    </div>
</template>

<script setup lang="ts">
import NotFound from "./NotFound.vue";
import ViewQuiz from "./ViewQuiz.vue";
defineOptions({
    routes: [
        { name: "Home", path: "/", component: ViewQuiz },
        { name: "NotFound", path: "/:pathMatch(.*)*", component: NotFound },
    ]
});
</script>

<style lang="scss" scoped>
@use "../assets/color.scss";

.app-body {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;

    color: var(--theme-text, #1a1a1a);
    background: var(--theme-bg, #f2e6d6);

    transition: background 160ms ease, color 160ms ease;
}
</style>
