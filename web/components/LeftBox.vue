<template>
  <div>
    <TitleCard title="More Pages" :list="list"/>
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
