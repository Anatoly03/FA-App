<template>
    <div class='app-body'>
        View Home
        <!-- <router-view></router-view> -->
    </div>
</template>

<script setup lang="ts">
import NotFound from "./NotFound.vue";
defineOptions({
    routes: [
        // { name: "Home", path: "/", component: ViewHome },
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
