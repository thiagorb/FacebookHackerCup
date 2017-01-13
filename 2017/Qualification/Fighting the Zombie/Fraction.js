class Fraction {
    constructor(numerator, denominator) {
        this.numerator = numerator;
        this.denominator = denominator;
    }

    addNumber(number) {
        this.numerator += number * this.denominator;

        this.normalize();
    }

    addFraction(fraction) {
        this.numerator = this.numerator * fraction.denominator + fraction.numerator * this.denominator;
        this.denominator = this.denominator * fraction.denominator;

        this.normalize();
    }

    multiplyFraction(fraction) {
        this.numerator *= fraction.numerator;
        this.denominator *= fraction.denominator;

        this.normalize();
    }

    divideNumber(number) {
        this.denominator *= number;

        this.normalize();
    }

    normalize() {
        if (this.numerator == 0) {
            this.denominator = 1;
        }

        if (this.numerator % this.denominator == 0) {
            this.numerator /= this.denominator;
            this.denominator = 1;
        }

        while ((this.numerator % 2 == 0) && (this.denominator % 2 == 0)) {
            this.numerator /= 2;
            this.denominator /= 2;
        }

        while ((this.numerator % 3 == 0) && (this.denominator % 3 == 0)) {
            this.numerator /= 3;
            this.denominator /= 3;
        }

        if (this.numerator == this.denominator) {
            this.numerator = 1;
            this.denominator = 1;
        }
    }

    evaluate() {
        return this.numerator / this.denominator;
    }
}

module.exports = Fraction;
