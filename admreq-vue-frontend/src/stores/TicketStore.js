import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useTicketStore = defineStore('TicketStore', {
    state: () => ({
        userTickets: [],
        departments: [],
        UM: false,
    }),
    actions: {

        async getUserTickets(token) {
            const res = await fetch(path+'tsr/tickets',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    mode: 'user',
                }),
            })
            const data = await res.json()
            if (res.status === 200) {
                console.log(data.tickets)
                this.userTickets = data.tickets
                this.UM = this.userTickets.filter((el) => el.unreadMessages === true).length > 0
            }
        },
        async createTicket(text, token, target_dep) {
            const res = await fetch(path+'tsr/create',{
                method: "POST",
                body: JSON.stringify({
                    text: text,
                    token: token,
                    target_dep: target_dep
                }),
            })
            if (res.status === 200) {
                this.getUserTickets(token)
            }
        },
        async getDepartments(token) {
            const res = await fetch(path+'departments',{
                method: "POST",
                body: JSON.stringify({
                    mode: "user",
                    token: token,
                })
            })
            const data = await res.json()
            if (res.status === 200) {
                this.departments = data.departments
            }
        },
    }
})