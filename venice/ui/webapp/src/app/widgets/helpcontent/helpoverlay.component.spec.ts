import {  ComponentFixture, TestBed } from '@angular/core/testing';
import { configureTestSuite } from 'ng-bullet';

import { HelpoverlayComponent } from './helpoverlay.component';
import { ControllerService } from '@app/services/controller.service';
import { ConfirmationService } from 'primeng/primeng';
import { SharedModule } from '@app/components/shared/shared.module';
import { MatIconRegistry } from '@angular/material';
import { LogPublishersService } from '@app/services/logging/log-publishers.service';
import { LogService } from '@app/services/logging/log.service';
import { RouterTestingModule } from '@angular/router/testing';
import { MessageService } from '@app/services/message.service';
import { Component, ViewChild } from '@angular/core';
import { TemplatePortalDirective, ComponentPortal, PortalModule } from '@angular/cdk/portal';
import { Overlay } from '@angular/cdk/overlay';
import { By } from '@angular/platform-browser';
import { TestingUtility } from '@app/common/TestingUtility';
import { MaterialdesignModule } from '@app/lib/materialdesign.module';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';

@Component({
  template: `<div id="testcontainer">component help</div>
             `,
})
class DummyHelpComponent {
}

 @Component({
  template: `
              <ng-template cdkPortal #helpTemplate="cdkPortal">
              <div>template help</div>
              </ng-template>
             `,
})
class DummyComponent {
  @ViewChild('helpTemplate') helpTemplate: TemplatePortalDirective;

  constructor(protected controllerService: ControllerService) {}

  setHelpTextComponent() {
    this.controllerService.setHelpOverlayData({
      component: new ComponentPortal(DummyHelpComponent)
    });
  }

  setHelpTextTemplate() {
    this.controllerService.setHelpOverlayData({
      template: this.helpTemplate
    });
  }

  removeHelpText() {
    this.controllerService.setHelpOverlayData(null);
  }

 }

describe('HelpoverlayComponent', () => {
  let component: HelpoverlayComponent;
  let dummyComponent: DummyComponent;
  let fixture: ComponentFixture<HelpoverlayComponent>;
  let dummyFixture: ComponentFixture<DummyComponent>;
  let tu: TestingUtility;

  configureTestSuite(() => {
     TestBed.configureTestingModule({
      declarations: [HelpoverlayComponent, DummyComponent, DummyHelpComponent],
      imports: [
        PortalModule,
        MaterialdesignModule,
        RouterTestingModule,
      ],
      providers: [
        ControllerService,
        ConfirmationService,
        LogService,
        LogPublishersService,
        MatIconRegistry,
        MessageService,
        Overlay
      ],
    }).overrideModule(BrowserDynamicTestingModule, { set: { entryComponents: [DummyHelpComponent] } });
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HelpoverlayComponent);
    component = fixture.componentInstance;
    dummyFixture = TestBed.createComponent(DummyComponent);
    dummyComponent = dummyFixture.componentInstance;
    tu = new TestingUtility(fixture);
    fixture.detectChanges();
    dummyFixture.detectChanges();
  });

  it('should load help content', () => {
    // Should start with default text
    const helpButton = fixture.debugElement.query(By.css('button'));
    tu.sendClick(helpButton);
    let defaultHelpContent = fixture.debugElement.query(By.css('.helpoverlay-content h1'));
    expect(defaultHelpContent).toBeTruthy();
    expect(defaultHelpContent.nativeElement.textContent).toContain('default help');

    // Should load component text
    dummyComponent.setHelpTextComponent();
    fixture.detectChanges();
    let helpContent = fixture.debugElement.query(By.css('.helpoverlay-content div'));
    expect(helpContent).toBeTruthy();
    expect(helpContent.nativeElement.textContent).toContain('component help');

    // Should load template text
    dummyComponent.setHelpTextTemplate();
    fixture.detectChanges();
    helpContent = fixture.debugElement.query(By.css('.helpoverlay-content div'));
    expect(helpContent).toBeTruthy();
    expect(helpContent.nativeElement.textContent).toContain('template help');

    // Should go back to default text
    dummyComponent.removeHelpText();
    fixture.detectChanges();
    defaultHelpContent = fixture.debugElement.query(By.css('.helpoverlay-content h1'));
    expect(defaultHelpContent).toBeTruthy();
    expect(defaultHelpContent.nativeElement.textContent).toContain('default help');
  });

});
