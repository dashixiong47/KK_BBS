// composables/useToRoute.js
import { useRouter } from "#app"

export default function useToRoute() {
    const localePath = useLocalePath();
    const router = useRouter()
    let to = (url) => {
        router.push(localePath(url))
    }
    return { to };
}
