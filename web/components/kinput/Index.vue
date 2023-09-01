<template>
  <input class="glass bg-[rgba(0,0,0,0)] primary-text" type="text" :value="modelValue" @input="setValue">
</template>

<script setup>
let { modelValue } = defineProps(["modelValue"]);
let emit = defineEmits(["update:modelValue"]);
const setValue = (e) => {
  emit("update:modelValue", e.target.value);
};
</script>
