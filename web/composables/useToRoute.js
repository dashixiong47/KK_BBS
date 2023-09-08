// composables/useToRoute.js
import { useRouter } from "#app"
const localePath = useLocalePath();
export default function useToRoute() {
    const router = useRouter()
    let to = (url) => {
        router.push(localePath(url))
    }
    return { to };
}
