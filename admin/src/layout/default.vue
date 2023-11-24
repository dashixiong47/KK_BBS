<template>
  <div class="main flex h-[100vh] w-full">
    <div class="left bg-[#545c64] h-full">
      <Menu v-model="isCollapse" />
    </div>
    <div class="right flex flex-1 flex-col">
      <Nav @switchCollapse="(val) => (isCollapse = val)"></Nav>
      <div class="content flex-1 w-full p-5">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import Nav from "@/components/Nav.vue";
import Menu from "@/components/Menu.vue";
import router from "../router";
const isCollapse = ref(false);
</script>

<style scoped lang="postcss"></style>
