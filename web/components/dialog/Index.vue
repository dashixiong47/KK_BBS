<template>
  <Teleport to="body">
    <div
      v-if="modelValue"
      @click.stop="setValue"
      class="flex items-center justify-center bg-[--masking-color] dark:bg-[--dark-masking-color] fixed top-0 left-0 w-screen h-screen z-[9999]"
    >
      <div class="p-5 shadow-center rounded-2xl" @click.stop>
        <p class="w-full flex justify-between">
          <span>{{ title }}</span>
          <Icon
            class="cursor-pointer"
            name="ep:close-bold"
            size="1.5rem"
            @click="setValue"
          />
        </p>
        <slot></slot>
        <div class="flex justify-end">
          <KButton v-if="cancelBtn" @click="cancel" class="primary-text mr-4">
            {{ cancelText }}
          </KButton>
          <KButton v-if="confirmBtn" @click="confirm" class="primary-text">
            {{ confirmText }}
          </KButton>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
// let { modelValue } = defineProps(["modelValue"]);
let {
  modelValue,
  close,
  title,
  cancelText,
  confirmText,
  cancelBtn,
  confirmBtn,
} = defineProps({
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
  cancelText: {
    type: String,
    default: "取消",
  },
  confirmText: {
    type: String,
    default: "确定",
  },
  cancelBtn: {
    type: Boolean,
    default: false,
  },
  confirmBtn: {
    type: Boolean,
    default: true,
  },
});
let emit = defineEmits(["update:modelValue", "cancel", "confirm"]);
const setValue = (e) => {
  emit("update:modelValue", false);
};
function cancel() {
  emit("cancel");
  emit("update:modelValue", false);
}
function confirm() {
  emit("confirm");
  emit("update:modelValue", false);
}
</script>
