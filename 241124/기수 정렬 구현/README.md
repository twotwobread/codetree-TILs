# [기수 정렬 구현 ](https://www.codetree.ai/missions/6/problems/implement-radix-sort)

|유형|문제 경험치|난이도|
|---|---|---|
|[Novice High / 정렬 / 기수 정렬](https://www.codetree.ai/missions?missionId=6)|30xp|![어려움][hard]|








[b5]: https://img.shields.io/badge/Bronze_5-%235D3E31.svg
[b4]: https://img.shields.io/badge/Bronze_4-%235D3E31.svg
[b3]: https://img.shields.io/badge/Bronze_3-%235D3E31.svg
[b2]: https://img.shields.io/badge/Bronze_2-%235D3E31.svg
[b1]: https://img.shields.io/badge/Bronze_1-%235D3E31.svg
[s5]: https://img.shields.io/badge/Silver_5-%23394960.svg
[s4]: https://img.shields.io/badge/Silver_4-%23394960.svg
[s3]: https://img.shields.io/badge/Silver_3-%23394960.svg
[s2]: https://img.shields.io/badge/Silver_2-%23394960.svg
[s1]: https://img.shields.io/badge/Silver_1-%23394960.svg
[g5]: https://img.shields.io/badge/Gold_5-%23FFC433.svg
[g4]: https://img.shields.io/badge/Gold_4-%23FFC433.svg
[g3]: https://img.shields.io/badge/Gold_3-%23FFC433.svg
[g2]: https://img.shields.io/badge/Gold_2-%23FFC433.svg
[g1]: https://img.shields.io/badge/Gold_1-%23FFC433.svg
[p5]: https://img.shields.io/badge/Platinum_5-%2376DDD8.svg
[p4]: https://img.shields.io/badge/Platinum_4-%2376DDD8.svg
[p3]: https://img.shields.io/badge/Platinum_3-%2376DDD8.svg
[p2]: https://img.shields.io/badge/Platinum_2-%2376DDD8.svg
[p1]: https://img.shields.io/badge/Platinum_1-%2376DDD8.svg
[passed]: https://img.shields.io/badge/Passed-%23009D27.svg
[failed]: https://img.shields.io/badge/Failed-%23D24D57.svg
[easy]: https://img.shields.io/badge/쉬움-%235cb85c.svg?for-the-badge
[medium]: https://img.shields.io/badge/보통-%23FFC433.svg?for-the-badge
[hard]: https://img.shields.io/badge/어려움-%23D24D57.svg?for-the-badge

# 코드 관련
- 처음에는 코멘트한 코드를 이용했음 -> append 때문에 시간적 효율성과 2개의 10N의 배열을 사용하니까 공간적 효율성도 떨어질 것이라 판단.  
- 그래서 counting sort를 이용하여 append를 없애고 N의 배열을 이용. -> 근데 시간적인 효율성이 나아지지 않았음. -> 그 이유는 append의 최적화 (한번 메모리 공간할당 시 2배씩 할당해서 자주 발생하지 않음.)라 판단했음.  
- 그러면 최대한 효율적으로 변경하기 위해서 메모리 할당을 최소화하고 copy하는 비용을 없애려고 함. -> 배열을 재사용하고 copy 대신 포인터를 변경해서 한 10ms 정도 줄였음.