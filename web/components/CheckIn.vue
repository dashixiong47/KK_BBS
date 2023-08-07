<template>
  <div>
    <p class="border-b py-2 mb-2 font-semibold">签到</p>

    <div class="flex justify-between border-b pb-2">
      <span class="text-sm flex items-center">
        <Icon name="icon-jinbi" class="mr-2" />
        1111
      </span>
      <button
        class="text-sm flex items-center border rounded-md px-5 py-1"
        :class="{
          'text-blue-400': checkInStatus,
          'border-blue-400': checkInStatus,
          'cursor-not-allowed': checkInStatus,
        }"
        :disabled="checkInStatus"
      >
        <Icon name="icon-qiandao" class="mr-2" />
        {{ checkInStatus ? "已签到" : "立即签到" }}
      </button>
    </div>
    <div class="flex my-4 pb-4 border-b">
      <div
        v-for="(step, index) in steps"
        :key="index"
        class="flex-1 text-center"
      >
        <div
          :class="[
            currentStep >= index + 1 ? 'bg-blue-500' : 'bg-gray-300',
            'w-5 h-5 rounded-full mx-auto',
          ]"
        >
          {{ index + 1 }}
        </div>
        <div
          v-if="index < steps.length - 1"
          class="flex-1 border-t -m-2 ml-10"
          :class="
            currentStep > index + 1 ? 'border-blue-500' : 'border-gray-300'
          "
        ></div>
      </div>
    </div>
    <p class="border-b pb-2 mb-2 font-semibold">签到排行榜</p>
    <ul>
      <li v-for="iten in 6" class="border-b flex items-center justify-between py-2">
        <div>
          头像
          <!-- <div>name</div> -->
        </div>
        <button class="text-sm flex items-center border rounded-md px-5 py-1">
          关注
        </button>
      </li>
    </ul>
  </div>
</template>

<script setup>
let checkInStatus = ref(true);
const steps = ref([1, 2, 3, 4, 5, 6, 7]);
const { currentStep } = defineProps({
  currentStep: {
    type: Number,
    default: 1,
  },
});
</script>
