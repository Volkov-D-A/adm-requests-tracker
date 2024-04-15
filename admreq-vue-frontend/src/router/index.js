import {createRouter, createWebHistory} from "vue-router"
import ticketsView from "../views/ticketsView.vue"
import usersView from "../views/usersView.vue"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "tickets",
            component: ticketsView
        },
        {
            path: "/users",
            name: "users",
            component: usersView
        },
    ]
})

export default router