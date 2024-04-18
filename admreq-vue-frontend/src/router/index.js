import {createRouter, createWebHistory} from "vue-router"
import ticketsView from "../views/ticketsView.vue"
import usersView from "../views/usersView.vue"
import employeeView from "../views/employeeView.vue"
import adminView from "../views/adminView.vue"

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
        {
            path: "/works",
            name: "works",
            component: employeeView
        },
        {
            path: "/admin",
            name: "admin",
            component: adminView
        },
    ]
})

export default router