import { HttpRequest, InvocationContext } from "@azure/functions";
import { tq } from "../src/functions/tq";
import {test, jest, expect} from '@jest/globals';
import { spawnSync, SpawnSyncReturns } from 'child_process';

jest.mock('child_process');
const spawnSyncMocked = jest.mocked(spawnSync);

test('tq gets verb, object and optional variant from the route', () => {
    var request = new HttpRequest({
        url: "https://some.url",
        method: "get",
        params: {
            "object":"object",
            "id":"id"
        }
    })

    spawnSyncMocked.mockReturnValue(<SpawnSyncReturns<string>>{
        pid: 0,
        output: null,
        stdout: "{}",
        stderr: "",
        status: 0,
        signal: null
    });

    tq(request, new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(1);
    expect(spawnSyncMocked.mock.calls[0][1]).toContain("GET");
    expect(spawnSyncMocked.mock.calls[0][1]).toContain("object");

    var request = new HttpRequest({
        url: "https://some.url",
        method: "post",
        params: {
            "object":"another_object",
            "variant":"flag",
            "id":"id"
        }
    })

    tq(request, new InvocationContext())
    expect(spawnSyncMocked.mock.calls).toHaveLength(2);
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("POST");
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("another_object");
    expect(spawnSyncMocked.mock.calls[1][1]).toContain("--flag");
    expect(spawnSyncMocked.mock.calls[1][2]["input"]).toBe('{"id":"id"}');
})