<template>
  <div>
    <p class="border-b py-2 mb-2 font-semibold font-main-color">签到</p>

    <div class="flex justify-between border-b pb-2 font-main-color">
      <span class="text-sm flex items-center">
        <Icon name="healthicons:coins" class="mr-2" size="1.5rem" />
        1111
      </span>
      <button
        class="glass before:rounded-md border text-sm flex items-center rounded-md px-5 py-1 "
        :class="{
          'color-primary': checkInStatus,
          'cursor-not-allowed': checkInStatus,
        }"
        :disabled="checkInStatus"
      >
        <Icon name="ph:calendar-check" class="mr-2" />
        {{ checkInStatus ? "已签到" : "立即签到" }}
      </button>
    </div>
    <div class="flex my-4">
      <div
        v-for="(step, index) in steps"
        :key="index"
        class="relative flex-1 text-center text-sm text-light-2 dark:text-light-1"
      >
        <div
          :class="[currentStep >= index + 1 ? 'bg-blue-500' : 'bg-gray-200']"
          class="w-5 h-5 rounded-full mx-auto z-10 relative"
        >
          {{ index + 1 }}
        </div>
        <div
          v-if="index < steps.length - 1"
          :class="[currentStep >= index + 1 ? 'bg-blue-500' : 'bg-gray-200']"
          class="absolute z-9 w-full h-1 left-1/2 top-1/2 -translate-y-1/2"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup>
let checkInStatus = ref(true);
const steps = ref([1, 2, 3, 4, 5, 6, 7]);
const { currentStep } = defineProps({
  currentStep: {
    type: Number,
    default: 2,
  },
});
</script>
