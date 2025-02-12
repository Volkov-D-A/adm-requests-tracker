import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useAdminStore = defineStore('AdminStore', {
    state: () => ({
        adminTickets: [],
        showTickets: [],
        UM: false,
    }),
    actions:{
        async getAdminTickets(token) {
            const res = await fetch(path+'tsr/tickets',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    mode: 'admin',
                }),
            })
            const data = await res.json()
            if (res.status === 200) {
                console.log(data.tickets)
                this.adminTickets = data.tickets
                this.showTickets = this.adminTickets
                this.UM = this.adminTickets.filter((el) => el.unreadMessages === true).length > 0
            }
        },
        filterAll() {
            this.showTickets = this.adminTickets
        },
        filterNotEmployee() {
            this.showTickets = this.adminTickets.filter((el) => el.employeeId === "")
        },
        filterFinished() {
            this.showTickets = this.adminTickets.filter((el) => el.finished)
        },
    }, 
})