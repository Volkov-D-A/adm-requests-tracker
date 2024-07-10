import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useTicketStore = defineStore('TicketStore', {
    state: () => ({
        userTickets: [],
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
            }
        },
        async createTicket(text, token) {
            const res = await fetch(path+'tsr/create',{
                method: "POST",
                body: JSON.stringify({
                    text: text,
                    token: token
                }),
            })
            if (res.status === 200) {
                this.getUserTickets(token)
            }
        },
    }
})