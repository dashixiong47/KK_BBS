<template>
  <div class="flex justify-center">
    <div class="flex items-center shadow-center rounded-2xl  overflow-hidden">
      <div
        v-for="(tab, index) in tabs"
        :key="index"
        :class="[
          'relative text-sm h-8 w-20 uncheckedColor',
          'cursor-pointer ml-[1px]',
          index === 0 ? '!ml-0' : '',
          modelValue === index ? 'activeColor' : '',
        ]"
        @click="actived(index)"
      >
        <!-- <div
          class="absolute inset-0 rounded-full shadow-md transform transition-all duration-300"
          :class="activeTab === index ? 'scale-100' : 'scale-0'"
        ></div> -->
        <div class="w-full absolute inset-0 flex items-center justify-center">
          {{ tab.name }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
let activeTab = ref(0);
/**
 * @type {import('vue').DefineProps<{
 * tabs: {
 * name: string;
 * type: string;
 * }[];
 * modelValue: number | boolean;
 *
 */
let props = defineProps({
  tabs: {
    type: Array,
    required: true,
  },
  modelValue: {
    type: [Number, Boolean],
    default: false,
  },
});
const emit = defineEmits(["update:modelValue", "selectd"]);
const tabs = ref(props.tabs);
const modelValue = computed({
  get: () => {
    if (props.modelValue) {
      return props.modelValue;
    }
    return activeTab.value;
  },
  set: (value) => {
    if (props.modelValue) {
      emit("update:modelValue", value);
      return;
    }
    activeTab.value = value;
  },
});

const actived = (index) => {
  modelValue.value = index;
  emit("selectd", index);
};
</script>
