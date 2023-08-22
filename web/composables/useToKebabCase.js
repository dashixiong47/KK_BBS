// composables/useToKebabCase.js
import { ref, computed } from 'vue';

export default function useToKebabCase() {
    function toKebabCase(str) {
        return str
            .replace(/([a-z0-9])([A-Z])/g, '$1-$2') // 在小写字母和大写字母之间插入连字符
            .toLowerCase(); // 转换为小写
    }

    return { toKebabCase };
}
