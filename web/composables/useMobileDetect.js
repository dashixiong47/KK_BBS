// composables/useMobileDetect.js
import { ref, computed } from 'vue';

export default function useMobileDetect() {
    const isMobile = ref(false);
    onMounted(() => {
        isMobile.value = window.innerWidth <= 768;
        window.addEventListener('resize', () => {
            isMobile.value = window.innerWidth <= 768;
        });
    });
    return { isMobile };
}
