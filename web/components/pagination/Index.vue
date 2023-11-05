<template>
  <div class="flex flex-col">
    <KButton
      class="w-full mb-5"
      :class="{
        '!bg-[--disable-color]':
          modelValue.currentPage === total,
      }"
      :disabled="modelValue.currentPage === total"
      @click="change(modelValue.currentPage + 1)"
    >
      下一页
    </KButton>
    <div class="w-full flex justify-end">
      <div class="flex">
        <!-- 上一页按钮 -->
        <button
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            '!bg-[--disable-color]': modelValue.currentPage === 1,
            'cursor-not-allowed': modelValue.currentPage === 1,
          }"
          :disabled="modelValue.currentPage === 1"
          @click="change(modelValue.currentPage - 1)"
        >
          <Icon name="iconoir:nav-arrow-left" />
        </button>

        <!-- 显示省略号 -->
        <button
          v-if="shouldShowEllipsis && modelValue.currentPage > 4"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          @click="moveLeft"
        >
          ...
        </button>
        <!-- 显示中间的页数按钮 -->
        <button
          v-for="item in pages"
          :key="item"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            activeColor: item === modelValue.currentPage,
          }"
          @click="change(item)"
        >
          {{ item }}
        </button>

        <!-- 显示省略号 -->
        <button
          v-if="shouldShowEllipsis && total - modelValue.currentPage >= 4"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          @click="moveRight"
        >
          ...
        </button>

        <!-- 下一页按钮 -->
        <button
          class="shadow-center uncheckedColor w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            '!bg-[--disable-color]': modelValue.currentPage === total,
            'cursor-not-allowed': modelValue.currentPage === total,
          }"
          :disabled="modelValue.currentPage === total"
          @click="change(modelValue.currentPage + 1)"
        >
          <Icon name="iconoir:nav-arrow-right" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, reactive } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      total: 0,
      pageSize: 10,
      currentPage: 1,
    }),
  },
});

const modelValue = reactive({
  ...props.modelValue,
});
const total = computed(() =>
  modelValue.total ? Math.ceil(modelValue.total / modelValue.pageSize) : 1
);
console.log(total);
const shouldShowEllipsis = computed(() => total.value > 9); // 根据总页数决定是否显示省略号
const emit = defineEmits(["update:modelValue"]);

function change(index) {
  modelValue.currentPage = index;
  emit("update:modelValue", modelValue);
}

function moveLeft() {
  change(Math.max(1, modelValue.currentPage - 5)); // 向左移动5页，但不超过第1页
}

function moveRight() {
  change(Math.min(total.value, modelValue.currentPage + 5)); // 向右移动5页，但不超过最后一页
}

const pages = computed(() => {
  let start, end;
  console.log(total.value, total.value <= 9);
  if (total.value <= 9) {
    start = 1;
    end = total.value;
  } else {
    const sideButtons = 3; // 当总页数大于9时，每边显示3个按钮
    start = modelValue.currentPage - sideButtons;
    end = modelValue.currentPage + sideButtons;

    if (start < 1) {
      end += 1 - start; // 如果起始页码小于1，则相应增加结束页码
      start = 1; // 确保起始页码不小于1
    }
    if (end > total.value) {
      start -= end - total.value; // 如果结束页码大于总页数，则相应减少起始页码
      end = total.value; // 确保结束页码不超过总页数
    }
  }
  return Array.from({ length: end - start + 1 }, (_, index) => start + index); // 生成页码数组
});
</script>
