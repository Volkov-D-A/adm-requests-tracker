import {createRouter, createWebHistory} from "vue-router"
import ticketsView from "../views/ticketsView.vue"
import usersView from "../views/usersView.vue"
import employeeView from "../views/employeeView.vue"
import adminView from "../views/adminView.vue"
import archiveView from "../views/archiveView.vue"
import statView from "../views/statView.vue"
import fullView from "../views/fullView.vue"

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
        {
            path: "/archive",
            name: "archive",
            component: archiveView
        },
        {
            path: "/stat",
            name: "stat",
            component: statView
        },
        {
            path: "/full/:id/:mode",
            name: "full",
            component: fullView,
            props: true
        },
    ]
})

export default router