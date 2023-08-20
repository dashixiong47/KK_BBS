<template>
  <div class="fixed right-5 bottom-5 z-50">
    <FloatingTheme :index="0" class="mb-5" />
    <!-- <FloatingTheme :index="1" /> -->
  </div>
</template>

<script setup>
import { ref, provide } from "vue";
let index = ref(0);
provide("actived", index);
</script>
