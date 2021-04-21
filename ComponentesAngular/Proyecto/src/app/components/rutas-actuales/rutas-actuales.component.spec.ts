import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RutasActualesComponent } from './rutas-actuales.component';

describe('RutasActualesComponent', () => {
  let component: RutasActualesComponent;
  let fixture: ComponentFixture<RutasActualesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RutasActualesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RutasActualesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
