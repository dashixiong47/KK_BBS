<template>
  <div>
    <label
      :class="[
        'shadow-center relative inline-block w-14 h-6 align-middle select-none transition-normal cursor-pointer  rounded-full transition-transform duration-300',
        { 'bg-[--color-primary]': isOn },
      ]"
      @click="toggle"
    >
      <span
        :class="[
          'absolute top-0 left-0 h-full w-6 bg-white border-2 border-gray-300 rounded-full shadow transition-transform duration-300',
          {
            'translate-x-0': !isOn,
            'translate-x-8': isOn,
          },
        ]"
      ></span>
    </label>
  </div>
</template>

<script setup>
const props = defineProps({
  modelValue: {
    type: [String, Number, Boolean],  // 允许不同类型的值
    default: false,
  },
  onValue: {
    type: [String, Number, Boolean],  // 允许不同类型的值
    default: true,
  },
  offValue: {
    type: [String, Number, Boolean],  // 允许不同类型的值
    default: false,
  },
});

const emit = defineEmits(["update:modelValue"]);

const isOn = computed(() => props.modelValue === props.onValue);  // 计算是否为 "on" 状态

const toggle = () => {
  const newValue = isOn.value ? props.offValue : props.onValue;  // 根据当前状态切换值
  emit("update:modelValue", newValue);
};
</script>
