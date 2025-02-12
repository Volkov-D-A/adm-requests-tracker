import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useEmployeeStore = defineStore('EmployeeStore', {
        state: () => ({
                employeeTickets: [],
                UM: false,
            }),
            actions: {
                async getEmployeeTickets(token) {
                    const res = await fetch(path+'tsr/tickets',{
                        method: "POST",
                        body: JSON.stringify({
                            token: token,
                            mode: 'employee',
                        }),
                    })
                    const data = await res.json()
                    if (res.status === 200) {
                        console.log(data.tickets)
                        this.employeeTickets = data.tickets
                        this.UM = this.employeeTickets.filter((el) => el.unreadMessages === true).length > 0
                    }
                },
            }

})