import {openDB} from 'idb/with-async-ittr';

class LocalDatabase {
    db: Promise<IDBDatabase>

    constructor() {
        if(process.server)
            return

        this.db = openDB('cougnat', 1, {
            upgrade(db) {
                db.createObjectStore('reports');
            },
        });
    }
    
    async put(collection: string, key: string, value: any): Promise<string> {
        const db = await this.db;
        await db.put(collection, value, key)
        return key
    }

    async get<T>(collection: string, key: string): Promise<T> {
        return (await this.db).get(collection, key)
    }

    async * cursor(collection: string): AsyncIterable<any> {
        const tx = (await this.db).transaction(collection);
        
        for await (const cursor of tx.store) {
            yield cursor.value
        }

    }
}


export default new LocalDatabase()
