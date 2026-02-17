import http from 'k6/http';

export const options = {
    scenarios: {
        constant_rate: {
            executor: 'constant-arrival-rate',
            rate: 2000,          // 👈 requests per second
            timeUnit: '1s',
            duration: '30s',   // 👈 test duration
            preAllocatedVUs: 20,
            maxVUs: 2000,
        },
    },
};

export default function () {
    http.get('http://146.190.191.97/daemon');
}
