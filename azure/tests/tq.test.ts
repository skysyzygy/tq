import { HttpRequest, InvocationContext } from "@azure/functions";
import { tq } from "../src/functions/tq";
import {test, jest, expect} from '@jest/globals';
import { spawnSync, SpawnSyncReturns } from 'child_process';

jest.mock('child_process');
const spawnSyncMocked = jest.mocked(spawnSync);
spawnSyncMocked.mockReturnValue(<SpawnSyncReturns<string>>{
    pid: 0,
    output: null,
    stdout: "{}",
    stderr: "",
    status: 0,
    signal: null
});


test('tq gets verb, object and optional variant from the route', () => {
    tq(new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        }
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(1);
    expect(spawnSyncMocked.mock.calls[0][1]).toContain("GET");
    expect(spawnSyncMocked.mock.calls[0][1]).toContain("object");

    tq(new HttpRequest({
        url: "https://some.url",
        method: "post",
        params: {
            "object":"another_object",
            "variant":"flag",
            "id":"id"
        }
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(2);
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("POST");
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("another_object");
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("--flag");
    expect(spawnSyncMocked.mock.calls[1][2]["input"]).toBe('{"id":"id"}');
})

test("tq gets TQ_LOGIN env from headers with fallback to the process environment",() => {
    spawnSyncMocked.mockClear()    
    process.env.TQ_LOGIN = "process_env";

    tq(new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        }
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(1);
    expect(spawnSyncMocked.mock.calls[0][2]["env"]["TQ_LOGIN"]).toBe("process_env")

    tq(new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        },
        headers: {"TQ_LOGIN": "http_request"}
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(2);
    expect(spawnSyncMocked.mock.calls[1][2]["env"]["TQ_LOGIN"]).toBe("http_request")
})

test("tq returns output or errors as JSON", async () => {
    spawnSyncMocked.mockClear();
    spawnSyncMocked.mockReturnValue(<SpawnSyncReturns<string>>{
        pid: 0,
        output: null,
        stdout: JSON.stringify({"output":"some stuff"}),
        stderr: "",
        status: 0,
        signal: null
    });

    var response = await tq(new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        }
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(1);
    expect(response.jsonBody).toStrictEqual({"output":"some stuff"});

    spawnSyncMocked.mockReturnValue(<SpawnSyncReturns<string>>{
        pid: 0,
        output: null,
        stdout: JSON.stringify({"output":"some stuff"}),
        stderr: "whoops!",
        status: -1,
        signal: null
    });

    var response = await tq(new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        }
    }), new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(2);
    expect(response.jsonBody).toStrictEqual({
        "message":"whoops!",
        "status":-1
    });
    expect(response.status).toBe(400);
})