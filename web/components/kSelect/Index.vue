<template>
  <div class="relative inline-block text-left w-40 h-11" ref="dropdown">
    <div>
      <button
        type="button"
        class="inline-flex justify-between w-full shadow-center rounded-xl px-4 py-3 text-sm font-medium text-gray-700 hover:bg-gray-50"
        @click="toggleDropdown"
      >
        {{ getSelected.name }}
        <Icon
          size="2rem"
          name="mdi:triangle-small-down"
          class="w-5 h-5 text-gray-400"
        />
      </button>
    </div>

    <div
      v-if="isOpen"
      class="z-10  origin-top-right absolute right-0 mt-2 w-full rounded-xl overflow-hidden shadow-center ring-1 ring-black ring-opacity-5 focus:outline-none"
    >
      <div
        class="py-1"
        role="menu"
        aria-orientation="vertical"
        aria-labelledby="menu-button"
        tabindex="-1"
      >
        <a
          v-for="item in options"
          :key="item.value"
          @click="selectItem(item)"
          class="text-gray-700 block px-4 py-2 text-sm hover:bg-gray-100"
          role="menuitem"
          tabindex="-1"
          >{{ item.name }}
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";
const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: "",
  },
  options: {
    type: Array,
    default: () => [],
  },
});
const emit = defineEmits(["update:modelValue","change"]);
const isOpen = ref(false);
const dropdown = ref(null);
const getSelected = computed(() => {
  return props.options.find((item) => item.value === props.modelValue) || {};
});
function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

function selectItem(item) {
  emit("update:modelValue", item.value);
  emit("change", item.value);
  isOpen.value = false;
}

function handleClickOutside(event) {
  if (dropdown.value && !dropdown.value.contains(event.target)) {
    isOpen.value = false;
  }
}

onMounted(() => {
  window.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  window.removeEventListener("click", handleClickOutside);
});
</script>
