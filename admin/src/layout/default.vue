<template>
  <el-container>
    <el-aside class="bg-[#00152a] h-[100vh] !w-auto absolute">
      <Menu v-model="isCollapse" />
    </el-aside>
    <el-container :class="getWidth" class="transition-all duration-300">
      <el-header>
        <Nav @switchCollapse="(val) => (isCollapse = val)"></Nav>
      </el-header>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed, ref } from "vue";
import Nav from "@/components/Nav.vue";
import Menu from "@/components/Menu.vue";
import router from "../router";
const isCollapse = ref(false);
const getWidth = computed(() => {
  return isCollapse.value ? "ml-[64px]" : "ml-[200px]";
});
</script>

<style scoped lang="postcss"></style>
