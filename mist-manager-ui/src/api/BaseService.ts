import axios from "axios";

class BaseService {
    protected _baseUrl: string;
    
    constructor(baseUrl: string) {
        this._baseUrl = baseUrl;
    }

    async get<T>(url: string): Promise<T> {
        const response = await axios.get<T>(`${this._baseUrl}${url}`);
        return response.data;
    }
    // other methods
}

export default BaseService;