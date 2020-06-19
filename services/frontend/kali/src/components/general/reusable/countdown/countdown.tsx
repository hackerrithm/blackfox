import React, { useEffect, useState } from "react";

const CountdownTimer = () => {
        const calculateTimeLeft = () => {
                const difference = +new Date("2020-10-10") - +new Date();
                let timeLeft: any = {};

                if (difference > 0) {
                        timeLeft = {
                                days: Math.floor(difference / (1000 * 60 * 60 * 24)),
                                hours: Math.floor((difference / (1000 * 60 * 60)) % 24),
                                minutes: Math.floor((difference / 1000 / 60) % 60),
                                seconds: Math.floor((difference / 1000) % 60)
                        };
                }

                return timeLeft;
        };

        const [timeLeft, setTimeLeft] = useState(calculateTimeLeft());

        useEffect(() => {
                setTimeout(() => {
                        setTimeLeft(calculateTimeLeft());
                }, 1000);
        });

        const timerComponents: any = [];

        Object.keys(timeLeft).forEach(interval => {
                if (!timeLeft[interval]) {
                        return;
                }

                timerComponents.push(
                        <span>
                                {timeLeft[interval]} {interval}{" "}
                        </span>
                );
        });

        return (
                <div>
                        {timerComponents.length ? timerComponents : <span>Time's up!</span>}
                </div>
        );
}

export default CountdownTimer;