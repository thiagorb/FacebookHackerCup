# Travel Restrictions
10 points

2020 has been a tough year for travelers. Air travel is especially problematic as passengers need to spend long periods in security lines, waiting areas, and crowded cabins where social distancing is difficult to maintain.

To minimize the spread of COVID-19, each airline has decided to reorganize all of their flight routes in a linear fashion. The airlines are hoping that by not making any one country a "hub", the spread of the virus will be significantly limited.

An airline's flights normally service NN countries, running in various directions. Amidst the pandemic, each airline has carefully arranged these NN countries in a sequence with each assigned a number from 11 to NN. Flights are limited to run only between pairs of countries that are adjacent in this sequence, with service in both directions. No other flights are available. That is, flights are available from country i to country j if and only if |i - j| = 1.

To make things more complicated, some countries have issued their own restrictions on incoming and outgoing travel. These restrictions are indicated by the characters I<sub>1..N</sub> and O<sub>1..N</sub>, each of which is either "N" or "Y":

- If I<sub>i</sub> = "N", then incoming flights to country ii from any other country are disallowed. Otherwise, if I<sub>i</sub> = "Y", they may be allowed.
- If O<sub>i</sub> = "N", then outgoing flights from country ii to any other country are disallowed. Otherwise, if O<sub>i</sub> = "Y", they may be allowed.

If a flight between adjacent countries is not disallowed by either the departure or arrival country's restrictions, then it's allowed.

As a consulting data scientist in the airline industry, your job is to determine which trips between the various countries are possible. Let P<sub>i, j</sub> = "Y" if it's possible to travel from country i to country j via a sequence of 0 or more flights (which may pass through other countries along the way), and P<sub>i, j</sub> = "N" otherwise. Note that P<sub>i, i</sub> is always "Y". Output this N*N matrix of characters.

## Input
Input begins with an integer T, the number of airlines. For each airline, there are three lines. The first line contains the integer N. The second line contains the length-N string I<sub>1..N</sub>. The third line contains the length-N string O<sub>1..N</sub>.

## Output
For the ith airline, output a line containing "Case #i:" followed by N more lines, the ith of which contains the length-N string P<sub>i, 1..N</sub>.

## Constraints
- 1 ≤ T ≤ 100
- 2 ≤ N ≤ 50

## Explanation of Sample
In the first case, there are two countries with no restrictions. Therefore, trips between all pairs of countries are possible.

In the second case, there are two countries, and traveling into country 1 is restricted. Since country 2 is the only country adjacent to country 1, the only impossible trip is from country 2 to country 1.

In the third case, there are two countries, both of which restrict inbound travel.

In the fourth case, one may not enter countries 2 or 3, nor exit country 4.

## Sample input

    5
    2
    YY
    YY
    2
    NY
    YY
    2
    NN
    YY
    5
    YNNYY
    YYYNY
    10
    NYYYNNYYYY
    YYNYYNYYNY


## Sample output

    Case #1: 
    YY
    YY
    Case #2: 
    YY
    NY
    Case #3: 
    YN
    NY
    Case #4: 
    YNNNN
    YYNNN
    NNYYN
    NNNYN
    NNNYY
    Case #5: 
    YYYNNNNNNN
    NYYNNNNNNN
    NNYNNNNNNN
    NNYYNNNNNN
    NNYYYNNNNN
    NNNNNYNNNN
    NNNNNNYYYN
    NNNNNNYYYN
    NNNNNNNNYN
    NNNNNNNNYY
