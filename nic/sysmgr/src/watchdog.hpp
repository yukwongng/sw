#ifndef __WATCHDOG_HPP__
#define __WATCHDOG_HPP__

#include <string>

#include "fault_watcher.hpp"
#include "switchroot_watcher.hpp"
#include "timer_watcher.hpp"

class AbstractWatchdog {
public:
    virtual void kick() = 0;
};
typedef std::shared_ptr<AbstractWatchdog> AbstractWatchdogPtr;

class Watchdog : public FaultReactor,
                 public SwitchrootReactor,
                 public TimerReactor
{
private:
    TimerWatcherPtr     timer_watcher;
    bool                kick_it;
    AbstractWatchdogPtr watchdog;
public:
    static std::shared_ptr<Watchdog> create();
    void stop();
    virtual void on_switchroot();
    virtual void on_fault(std::string reason);
    virtual void on_timer();
};
typedef std::shared_ptr<Watchdog> WatchdogPtr;

#endif
