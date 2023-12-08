<template>
  <div class="flex items-center">
    <KButton class="flex-shrink-0" @click="scrollLeft">
      <Icon
        class="w-6 h-6"
        name="material-symbols-light:keyboard-double-arrow-left"
      ></Icon>
    </KButton>
    <ul
      ref="scrollContainer"
      class="flex items-center mx-3 py-3 overflow-scroll hide-scrollbar smooth-scroll"
    >
      <li v-for="item in getGroup" :key="item.id" class="mx-2 flex-shrink-0">
        <KButton @click="change(item)" :actived="modelValue === item.id">{{
          item.name
        }}</KButton>
      </li>
    </ul>
    <KButton class="flex-shrink-0" @click="scrollRight">
      <Icon
        class="w-6 h-6"
        name="material-symbols-light:keyboard-double-arrow-right"
      ></Icon>
    </KButton>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useGroupStore } from "@/stores/init";
const groupStore = useGroupStore();
const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: null,
  },
});

let getGroup = computed(() =>
  groupStore.getGroup.filter((item) => item.id != undefined)
);

let emit = defineEmits(["update:modelValue", "change"]);
const scrollContainer = ref(null);

function change(item) {
  emit("update:modelValue", item.id);
  emit("change", item.id);
  groupStore.setActived(item.id);
}

function scrollLeft() {
  if (scrollContainer.value) {
    scrollContainer.value.scrollLeft -= scrollContainer.value.clientWidth / 2;
  }
}

function scrollRight() {
  if (scrollContainer.value) {
    scrollContainer.value.scrollLeft += scrollContainer.value.clientWidth / 2;
  }
}
// 获取小组列表
groupStore.fetchGroup();
</script>

<style scoped>
.hide-scrollbar::-webkit-scrollbar {
  display: none;
}

.hide-scrollbar {
  scrollbar-width: none;
}

.smooth-scroll {
  scroll-behavior: smooth;
}
</style>
