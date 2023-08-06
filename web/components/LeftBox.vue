<template>
  <div>
    <CardTitleCard title="More Pages" :list="list" class="mb-5 h-80 "/>
    <CardTitleCard title="More Pages" :list="list" class="mb-5 h-80 "/>
  </div>
</template>

<script setup>
// import { useCounterStore } from "../stores/main.js";
// const counter = useCounterStore();
// counter.increment();
// console.log(counter.double);
// console.log(counter.count);
const list = [0, 1, 2, 3, 4, 0, 1, 2, 3, 4];
</script>
