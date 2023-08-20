<template>
  <div style="color: red">
    {{ num }}
    <button @click="add">按钮</button>
  </div>
</template>

<script setup>
let num = ref(0);
function add() {
  console.log("------------");
  num.value++;
}
</script>
