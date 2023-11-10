<template>
  <input
    class="glass bg-[rgba(0,0,0,0)] font-main-color outline-none ring-0"
    :type="type"
    :value="displayValue"
    @input="setValue"
  />
</template>

<script setup>
let props = defineProps({
  modelValue: String | Number,
  modelModifiers: { default: () => ({}) },
  max: { default: false, type: [Number, Boolean] },
  min: { default: false, type: [Number, Boolean] },  // 定义 min 属性
  type: { default: "text", type: String },  
});
let emit = defineEmits(["update:modelValue"]);

const inputValue = ref(props.modelValue);

// 更新 inputValue 当 props.modelValue 改变
watchEffect(() => {
  inputValue.value = props.modelValue;
});

const displayValue = computed(() => {
  return props.modelModifiers.number
    ? String(clampValue(extractNumbers(inputValue.value)))
    : inputValue.value;
});

const setValue = (e) => {
  inputValue.value = e.target.value; // 更新 inputValue 的值
  if (props.modelModifiers.number) {
    let numericValue = extractNumbers(inputValue.value);
    numericValue = clampValue(numericValue);
    inputValue.value = numericValue; // 重新设置 inputValue 的值，如果它超过了最大/最小值
    emit("update:modelValue", numericValue); // 发送处理后的数字
  } else {
    emit("update:modelValue", inputValue.value); // 发送原始输入
  }
};

function extractNumbers(str) {
  if(str === '') return '';
  str = String(str); // 确保 str 是一个字符串
  return Number(str.replace(/\D/g, ""));
}

function clampValue(num) {
  if (props.min && num < props.min) {
    return props.min;  // 如果 num 小于 min，返回 min
  }
  if (props.max && num > props.max) {
    return props.max;  // 如果 num 大于 max，返回 max
  }
  return num;
}
</script>
