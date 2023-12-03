<template>
  <div class="relative inline-block text-left w-40" ref="dropdown">
    <div>
      <button
        type="button"
        class="inline-flex justify-between w-full shadow-center rounded-xl px-4 py-3 text-sm font-medium text-gray-700 hover:bg-gray-50"
        @click="toggleDropdown"
      >
        {{ selected }}
        <Icon
          size="2rem"
          name="mdi:triangle-small-down"
          class="w-5 h-5 text-gray-400"
        />
      </button>
    </div>

    <div
      v-if="isOpen"
      class="origin-top-right absolute right-0 mt-2 w-56 rounded-xl overflow-hidden shadow-center ring-1 ring-black ring-opacity-5 focus:outline-none"
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
          :key="item"
          @click="selectItem(item)"
          class="text-gray-700 block px-4 py-2 text-sm hover:bg-gray-100"
          role="menuitem"
          tabindex="-1"
          >{{ item }}
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

const isOpen = ref(false);
const selected = ref("请选择");
const options = ["选项1", "选项2", "选项3"];
const dropdown = ref(null);

function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

function selectItem(item) {
  selected.value = item;
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
