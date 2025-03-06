/* damageCalculation.js */

// 将伤害乘区计算方法添加到 Initialization 类原型

Initialization.prototype.basicDamageArea = function(mag) {
    this.output.basicDamageArea = this.currentPanel.attack * mag.magnificationValue / 100 * mag.triggerTimes;
};

Initialization.prototype.increasedDamageArea = function(mag) {
    this.output.increasedDamageArea = 1 + ((mag.increasedDamage || 0) + this.currentPanel.increasedDamage) / 100;
};

Initialization.prototype.explosiveInjuryArea = function() {
    this.output.explosiveInjuryArea = 1 + (this.currentPanel.critical * this.currentPanel.explosiveInjury) / 10000;
};

Initialization.prototype.defenseArea = function(mag) {
    const characterBase = 793.783;
    const totalDefense = 873.1613;
    let penetration = (this.defense.penetration - (mag.penetration || 0)) / 100;
    let defenseBreak = (this.defense.defenseBreak - (mag.defenseBreak || 0)) / 100;
    this.output.defenseArea = characterBase / (totalDefense * (1 - penetration) * (1 - defenseBreak) - this.defense.penetrationValue + characterBase);
};

Initialization.prototype.reductionResistanceArea = function(mag) {
    this.output.reductionResistanceArea = 1 + ((mag.reductionResistance || 0) + this.currentPanel.reductionResistance) / 100;
};

Initialization.prototype.vulnerableArea = function() {
    this.output.vulnerableArea = 1 + (this.currentPanel.vulnerable) / 100;
};

Initialization.prototype.specialDamageArea = function() {
    this.output.specialDamageArea = 1 + (this.currentPanel.specialDamage) / 100;
};

Initialization.prototype.initializationArea = function(mag) {
    this.basicDamageArea(mag);
    this.increasedDamageArea(mag);
    this.explosiveInjuryArea();
    this.defenseArea(mag);
    this.reductionResistanceArea(mag);
    this.vulnerableArea();
    this.specialDamageArea();
};

Initialization.prototype.calculatingTotalDamage = function() {
    let totalDamage = 0;
    for (let mag of this.magnifications) {
        this.initializationArea(mag);
        let damage = this.output.basicDamageArea *
            this.output.increasedDamageArea *
            this.output.explosiveInjuryArea *
            this.output.defenseArea *
            this.output.reductionResistanceArea *
            this.output.vulnerableArea *
            this.output.specialDamageArea *
            (1 + (mag.specialDamage || 0) / 100);
        totalDamage += damage;
    }
    return totalDamage;
};
