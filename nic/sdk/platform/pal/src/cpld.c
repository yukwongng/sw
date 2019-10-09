
/*
 * Copyright (c) 2018, Pensando Systems Inc.
 */
#include "pal_cpld.h"
#include "pal_locks.h"
#include "pal.h"

#ifdef __x86_64__
int pal_is_qsfp_port_psnt(int port_no)
{
    return 0;
}

int
pal_qsfp_set_port(int port)
{
    return -1;
}

int
pal_qsfp_reset_port(int port)
{
    return -1;
}

int
pal_qsfp_set_low_power_mode(int port)
{
    return -1;
}

int
pal_qsfp_reset_low_power_mode(int port)
{
    return -1;
}

int
pal_qsfp_set_led(int port, pal_led_color_t led,
                 pal_led_frequency_t frequency)
{
    return -1;
}

int
pal_program_marvell(uint8_t marvell_addr, uint32_t data)
{
    return -1;
}

int
pal_get_cpld_rev(void)
{
    return -1;
}

int
pal_marvell_link_status(uint8_t marvell_addr, uint16_t *data, uint8_t phy)
{
    return -1;
}

int
pal_get_cpld_id(void)
{
    return -1;
}

int
pal_system_set_led(pal_led_color_t led, pal_led_frequency_t frequency)
{
    return -1;
}

void
pal_write_core_temp(int data)
{
}

void
pal_write_hbm_temp(int data)
{
}

void
pal_write_board_temp(int data)
{
}

void
pal_cpld_reload_reset(void)
{
}

bool
pal_cpld_verify_idcode(void)
{
    return false;
}

int
pal_cpld_erase(void)
{
    return -1;
}

int
pal_cpld_read_flash(uint8_t *buf, uint32_t size)
{
    return -1;
}

int
pal_cpld_write_flash(const uint8_t *buf, uint32_t size, cpld_upgrade_status_cb_t cpld_upgrade_status_cb, void *arg)
{
    return -1;
}

void
pal_power_cycle(void)
{
}

bool
pal_cpld_hwlock_enabled(void)
{
    return false;
}
#else
#include <string.h>

static int CPLD_FAIL    = -1;
static int CPLD_SUCCESS = 0;
static int cpld_rev     = -1;
static int cpld_id      = -1;

/* Public APIs */
int
pal_is_qsfp_port_psnt(int port_no)
{
    int cpld_rd_data = cpld_reg_rd(CPLD_REGISTER_QSFP_CTRL);

    if (port_no == 1) {
        return ((cpld_rd_data & 0x10) != 0);
    } else if (port_no == 2) {
        return ((cpld_rd_data & 0x20) != 0);
    }
    return CPLD_FAIL;
}

/* Register: 0x2, bit 0 is qsfp port 1, bit 1 is qsfp port 2 */
int
pal_qsfp_set_port(int port)
{
    int bit = 0;

    if (port == 1) {
        bit = 0;
    } else if (port == 2) {
        bit = 1;
    } else {
        return CPLD_FAIL;
    }
    return cpld_reg_bit_reset(CPLD_REGISTER_QSFP_CTRL, bit);
}

int
pal_qsfp_reset_port(int port)
{
    int bit = 0;

    if (port == 1) {
        bit = 0;
    } else if (port == 2) {
        bit = 1;
    } else {
        return CPLD_FAIL;
    }
    return cpld_reg_bit_set(CPLD_REGISTER_QSFP_CTRL, bit);
}

/* Register: 0x2, bit 2 is qsfp port 1, bit 3 is qsfp port 2 */
int
pal_qsfp_set_low_power_mode(int port)
{
    int bit = 0;

    if (port == 1) {
        bit = 2;
    } else if (port == 2) {
        bit = 3;
    } else {
        return CPLD_FAIL;
    }
    return cpld_reg_bit_set(CPLD_REGISTER_QSFP_CTRL, bit);
}

int
pal_qsfp_reset_low_power_mode(int port)
{
    int bit = 0;

    if (port == 1) {
        bit = 2;
    } else if (port == 2) {
        bit = 3;
    } else {
        return CPLD_FAIL;
    }
    return cpld_reg_bit_reset(CPLD_REGISTER_QSFP_CTRL, bit);
}

static int
pal_change_qsfp_frequency(uint8_t mask, uint8_t frequency)
{
    uint8_t frequency_orig;
    frequency_orig = cpld_reg_rd(CPLD_REGISTER_QSFP_LED_FREQUENCY);
    frequency_orig = frequency_orig & mask;
    frequency_orig = frequency_orig | frequency;
    return cpld_reg_wr(CPLD_REGISTER_QSFP_LED_FREQUENCY, frequency_orig);
}

int
pal_qsfp_set_led(int port, pal_led_color_t led,
                 pal_led_frequency_t frequency)
{
    static uint8_t qsfp_port1_led_color;
    static uint8_t qsfp_port2_led_color;
    static uint8_t qsfp_port1_green_led_frequency;
    static uint8_t qsfp_port1_yellow_led_frequency;
    static uint8_t qsfp_port2_green_led_frequency;
    static uint8_t qsfp_port2_yellow_led_frequency;

    switch(port) {
        case 1:
            //check if the port1 LED has changed color
            if (qsfp_port1_led_color == led) {

                //check if frequency needs to be changed.
                if (led == LED_COLOR_GREEN &&
                    qsfp_port1_green_led_frequency != frequency) {
                        qsfp_port1_green_led_frequency =  frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT1_GREEN_BLINK_MASK,
                                                         frequency << QSFP_PORT1_GREEN_BLINK_SHIFT);
                } else if (led == LED_COLOR_YELLOW &&
                           qsfp_port1_yellow_led_frequency != frequency) {
                        qsfp_port1_yellow_led_frequency =  frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT1_YELLOW_BLINK_MASK,
                                                         frequency << QSFP_PORT1_YELLOW_BLINK_SHIFT);
                } else {
                    return CPLD_FAIL;
                }

            } else {
                qsfp_port1_led_color = led;
                switch(led) {
                    case LED_COLOR_GREEN:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_YELLOW_ON);
                        cpld_reg_bit_set(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_GREEN_ON);
                        qsfp_port1_green_led_frequency = frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT1_GREEN_BLINK_MASK,
                                                         frequency << QSFP_PORT1_GREEN_BLINK_SHIFT);
                    case LED_COLOR_YELLOW:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_GREEN_ON);
                        cpld_reg_bit_set(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_YELLOW_ON);
                        qsfp_port1_yellow_led_frequency = frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT1_YELLOW_BLINK_MASK,
                                                         frequency << QSFP_PORT1_YELLOW_BLINK_SHIFT);
                    case LED_COLOR_NONE:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_GREEN_ON);
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT1_YELLOW_ON);
                        return CPLD_SUCCESS;
                    default:
                        return CPLD_FAIL;
                }
            }
        case 2:
            //check if the port2 LED has changed color
            if (qsfp_port2_led_color == led) {

                //check if frequency needs to be changed.
                if (led == LED_COLOR_GREEN &&
                    qsfp_port2_green_led_frequency != frequency) {
                        qsfp_port2_green_led_frequency =  frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT2_GREEN_BLINK_MASK,
                                                         frequency << QSFP_PORT2_GREEN_BLINK_SHIFT);
                } else if (led == LED_COLOR_YELLOW &&
                           qsfp_port2_yellow_led_frequency != frequency) {
                        qsfp_port2_yellow_led_frequency =  frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT2_YELLOW_BLINK_MASK,
                                                         frequency << QSFP_PORT2_YELLOW_BLINK_SHIFT);
                } else {
                    return CPLD_FAIL;
                }

            } else {
                qsfp_port2_led_color = led;
                switch(led) {
                    case LED_COLOR_GREEN:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_YELLOW_ON);
                        cpld_reg_bit_set(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_GREEN_ON);
                        qsfp_port2_green_led_frequency = frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT2_GREEN_BLINK_MASK,
                                                         frequency << QSFP_PORT2_GREEN_BLINK_SHIFT);
                    case LED_COLOR_YELLOW:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_GREEN_ON);
                        cpld_reg_bit_set(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_YELLOW_ON);
                        qsfp_port2_yellow_led_frequency = frequency;
                        return pal_change_qsfp_frequency((uint8_t)QSFP_PORT2_YELLOW_BLINK_MASK,
                                                         frequency << QSFP_PORT2_YELLOW_BLINK_SHIFT);
                    case LED_COLOR_NONE:
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_GREEN_ON);
                        cpld_reg_bit_reset(CPLD_REGISTER_QSFP_LED, QSFP_LED_PORT2_YELLOW_ON);
                        return CPLD_SUCCESS;
                    default:
                        return CPLD_FAIL;
                }
            }
       default:
           return CPLD_FAIL;
    }
}

int
pal_program_marvell(uint8_t marvell_addr, uint32_t data)
{
    cpld_reg_wr(0x7, marvell_addr);
    cpld_reg_wr(0x8, (data >> 8) && 0xff);
    cpld_reg_wr(0x9, data && 0xff);
    cpld_reg_wr(0x6, (0xc << 3) | 0x4 | 0x1);
    cpld_reg_wr(0x6, 0);

    return CPLD_SUCCESS;
}

int
pal_marvell_link_status(uint8_t marvell_addr, uint16_t *data, uint8_t phy)
{
    int rc = CPLD_FAIL;
    rc = cpld_mdio_rd(marvell_addr, data, phy);
    *data =  (*data >> 2) & 0x1;
    return rc;
}

int
pal_get_cpld_rev(void)
{
    if (cpld_rev == -1) {
        cpld_rev = cpld_reg_rd(CPLD_REGISTER_REVISION);
    }
    return cpld_rev;
}

int
pal_get_cpld_id(void)
{
    if (cpld_id == -1) {
        cpld_id = cpld_reg_rd(CPLD_REGISTER_ID);
    }
    return cpld_id;
}

bool
pal_cpld_hwlock_enabled(void)
{
    int value = 0;

    value = cpld_reg_rd(CPLD_PERSISTENT_REG);
    if (value == -1) {
        return value;
    }
    return (value & CPLD_HWLOCK_MASK) ? true : false;
}

int
pal_system_set_led(pal_led_color_t led, pal_led_frequency_t frequency)
{
    switch(led) {
        case LED_COLOR_GREEN:
            switch(frequency) {
                case LED_FREQUENCY_0HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x14);
                case LED_FREQUENCY_2HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x15);
                case LED_FREQUENCY_1HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x16);
                case LED_FREQUENCY_05HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x17);
                default:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x00);
            }

        case LED_COLOR_YELLOW:
            switch(frequency) {
                case LED_FREQUENCY_0HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x20);
                case LED_FREQUENCY_2HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x28);
                case LED_FREQUENCY_1HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x30);
                case LED_FREQUENCY_05HZ:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x38);
                default:
                    return cpld_reg_wr(CPLD_REGISTER_SYSTEM_LED, 0x00);
            }
       default:
           return CPLD_FAIL;
    }
}

void
pal_write_core_temp(int data)
{
    data = (data > 127) ? 127 : data;
    cpld_reg_wr(CPLD_REGISTER_CORE_TEMP, data);
}

void
pal_write_hbm_temp(int data)
{
    data = (data > 127) ? 127 : data;
    cpld_reg_wr(CPLD_REGISTER_HBM_TEMP, data);
}

void
pal_write_board_temp(int data)
{
    data = (data > 127) ? 127 : data;
    cpld_reg_wr(CPLD_REGISTER_BOARD_TEMP, data);
}

void
pal_cpld_reload_reset(void)
{
    cpld_reload_reset();
}

bool
pal_cpld_verify_idcode(void)
{
    return cpld_verify_idcode();
}

int
pal_cpld_erase(void)
{
    return cpld_erase();
}

int
pal_cpld_read_flash(uint8_t *buf, uint32_t size)
{
    return cpld_read_flash(buf, size);
}

int
pal_cpld_write_flash(const uint8_t *buf, uint32_t size, cpld_upgrade_status_cb_t cpld_upgrade_status_cb, void *arg)
{
    return cpld_write_flash(buf, size, cpld_upgrade_status_cb, arg);
}

void
pal_power_cycle(void)
{
    pal_write_gpios(GPIO_3_POWER_CYCLE, GPIO_EN);
}
#endif
