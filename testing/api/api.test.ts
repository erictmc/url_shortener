import axios from "axios";
import each from "jest-each";


const apiUrl = "http://localhost:8080"

const createUrlEntry = (originalUrl: string) => {
    return axios.post(apiUrl + "/new", { original_url: originalUrl })
}

describe("valid api is working correctly",  () => {
    it("api correctly handles basic url and results in redirect",  async () => {
        const testUrl = "https://www.google.com/"
        const { data } = await createUrlEntry(testUrl)
        const { short_url, original_url } = data
        expect(original_url).toEqual(testUrl)
        const response = await axios.get(short_url)

        // Somewhat hacky code to detect redirect
        // https://stackoverflow.com/questions/55926127/does-axios-have-the-ability-to-detect-redirects
        expect(response.request._redirectable._redirectCount).toEqual(1)
        expect(response.status).toEqual(200);
        expect(response.request.res.responseUrl).toEqual(original_url)
    })

    describe("test invalid blank entries", () => {
        each(["", null, undefined ]).test("test case: %s", async (t) => {
            try {
                await createUrlEntry(t);
            } catch (e){
                expect(e.response.status).toEqual(422);
                expect(e.response.data).toEqual( { original_url: 'cannot be blank' })
            }
        })

        each([
            "foobar",
            "https://google.com more stuff",
            "www.google.com"
        ]).test("test case: %s", async (t) => {
            try {
                await createUrlEntry(t);
            } catch (e){
                expect(e.response.status).toEqual(422);
            }
        })
    })
})