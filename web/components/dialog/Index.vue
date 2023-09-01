<template>
  <Teleport to="body">
    <div
      v-if="modelValue"
      @click.stop="setValue"
      class="flex items-center justify-center bg-[rgba(0,0,0,0.2)] fixed top-0 left-0 w-screen h-screen z-[9999]"
    >
      <div class="p-5 bg-white  shadow-center rounded-2xl" @click.stop>
        <p class="w-full flex justify-between">
          <span>{{ title }}</span>
          <Icon class="cursor-pointer" name="ep:close-bold" size="1.5rem" @click="setValue" />
        </p>
        <slot></slot>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
// let { modelValue } = defineProps(["modelValue"]);
let { modelValue, close, title } = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    default: "提示",
  },
  close: {
    type: Boolean,
    default: true,
  },
});
let emit = defineEmits(["update:modelValue"]);
const setValue = (e) => {
  emit("update:modelValue", false);
};
</script>
